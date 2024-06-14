package vodka

import (
	"log/slog"
	"net/http"
	"reflect"
	"runtime"
)

var HttpErrorHandler func(err error, c Context)

func init() {
	HttpErrorHandler = func(err error, c Context) {
		slog.Error("Internal Server Error", err)

		c.String(http.StatusInternalServerError, "Internal Server Error")
	}
}

type Vodka struct {
	nextFn map[string]HandlerFunc
	route  string

	http.ServeMux
}

func New() *Vodka {
	return &Vodka{
		nextFn: map[string]HandlerFunc{},
		route:  "",
	}
}

func (v *Vodka) Static(pattern, directory string) {
	v.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(directory))))
}

func (v *Vodka) Use(next ...HandlerFunc) {
	for _, nextFn := range next {
		funcName := runtime.FuncForPC(reflect.ValueOf(nextFn).Pointer()).Name()
		if _, ok := v.nextFn[funcName]; !ok {
			v.nextFn[funcName] = nextFn
		}
	}
}

func (v *Vodka) GET(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodGet, v.route+pattern, handler, nextFn...)
}

func (v *Vodka) POST(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodPost, v.route+pattern, handler, nextFn...)
}

func (v *Vodka) PUT(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodPut, v.route+pattern, handler, nextFn...)
}

func (v *Vodka) DELETE(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodDelete, v.route+pattern, handler, nextFn...)
}

func (v *Vodka) PATCH(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodPatch, v.route+pattern, handler, nextFn...)
}

func (v *Vodka) OPTIONS(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodOptions, v.route+pattern, handler, nextFn...)
}

func (v *Vodka) HEAD(pattern string, handler HandlerFunc, nextFn ...HandlerFunc) {
	v.httpHandler(http.MethodHead, v.route+pattern, handler, nextFn...)
}

func (v *Vodka) Route(pattern string) *Vodka {
	if pattern[0] != '/' {
		panic("Route must start with /")
	}

	if pattern[len(pattern)-1] == '/' {
		pattern = pattern[:len(pattern)-1]
	}

	v.route = pattern

	return v
}

func (v *Vodka) Start(port string) error {
	return http.ListenAndServe(port, v)
}

func (v *Vodka) httpHandler(
	method, pattern string,
	handler HandlerFunc,
	nextFn ...HandlerFunc,
) {
	v.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			c := &ctx{
				w:      w,
				r:      r,
				url:    r.URL,
				method: r.Method,
				body:   r.Body,
			}

			for _, fn := range v.nextFn {
				if err := fn(c); err != nil {
					HttpErrorHandler(err, c)
					return
				}
			}

			for _, fn := range nextFn {
				if err := fn(c); err != nil {
					HttpErrorHandler(err, c)
					return
				}
			}

			handler(c)
		}
	})
}
