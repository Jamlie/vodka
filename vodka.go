package vodka

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"runtime"
)

// type HandlerFunc func(http.ResponseWriter, *http.Request, *Vodka)

type HandlerFunc func(c *Context)

type Vodka struct {
	nextFn map[string]HandlerFunc

	http.ServeMux
}

func New() *Vodka {
	return &Vodka{
		nextFn: map[string]HandlerFunc{},
	}
}

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func Wrap(fn http.HandlerFunc) HandlerFunc {
	return func(c *Context) {
		fn(c.w, c.r)
	}
}

func (c *Context) Request() *http.Request {
	return c.r
}

func (c *Context) Response() http.ResponseWriter {
	return c.w
}

func (c *Context) Context() context.Context {
	return c.r.Context()
}

func (c *Context) Query(query string) string {
	return c.Request().PathValue(query)
}

func (v *Vodka) Static(pattern, directory string) {
	v.Handle(
		pattern,
		http.StripPrefix(pattern, http.FileServer(http.Dir(directory))),
	)
}

func (v *Vodka) Use(next ...HandlerFunc) {
	for _, nextFn := range next {
		funcName := runtime.FuncForPC(reflect.ValueOf(nextFn).Pointer()).Name()
		if _, ok := v.nextFn[funcName]; !ok {
			v.nextFn[funcName] = nextFn
		}
	}
}

func (c *Context) XML(status int, data any) error {
	w := bytes.NewBuffer(make([]byte, 0))
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(status)
	_, err = w.WriteTo(c.w)
	return err
}

func (c *Context) JSON(status int, data any) error {
	w := bytes.NewBuffer(make([]byte, 0))
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(status)
	_, err = w.WriteTo(c.w)
	return err
}

func (c *Context) File(path string) {
	http.ServeFile(c.Response(), c.Request(), path)
}

func (c *Context) String(status int, data string) {
	b := bytes.NewBuffer([]byte(data))
	c.w.Header().Set("Content-Type", "text/plain")
	c.w.WriteHeader(status)
	_, _ = b.WriteTo(c.w)
}

func (c *Context) HTML(status int, data string) {
	b := bytes.NewBuffer([]byte(data))
	c.w.Header().Set("Content-Type", "text/html")
	c.w.WriteHeader(status)
	_, _ = b.WriteTo(c.w)
}

func (v *Vodka) httpHandler(
	method, pattern string,
	handler HandlerFunc,
	nextFn ...HandlerFunc,
) {
	v.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			c := &Context{
				w: w,
				r: r,
			}

			for _, fn := range v.nextFn {
				fn(c)
			}

			for _, fn := range nextFn {
				fn(c)
			}

			handler(c)
		}
	})
}

func (v *Vodka) GET(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodGet, pattern, handler, nextFn...)
}

func (v *Vodka) POST(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodPost, pattern, handler, nextFn...)
}

func (v *Vodka) PUT(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodPut, pattern, handler, nextFn...)
}

func (v *Vodka) DELETE(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodDelete, pattern, handler, nextFn...)
}

func (v *Vodka) PATCH(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodPatch, pattern, handler, nextFn...)
}

func (v *Vodka) OPTIONS(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodOptions, pattern, handler, nextFn...)
}

func (v *Vodka) HEAD(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodHead, pattern, handler, nextFn...)
}

func (v *Vodka) Start(port string) error {
	return http.ListenAndServe(port, v)
}
