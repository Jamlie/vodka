package vodka

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Context interface {
	Request() *http.Request
	Response() http.ResponseWriter
	Context() context.Context
	Query(string) string
	XML(int, any) error
	JSON(int, any) error
	File(string)
	String(int, string)
	HTML(int, string)
}

type HandlerFunc func(c Context)

type defaultContext struct {
	w http.ResponseWriter
	r *http.Request
}

func Wrap(fn http.HandlerFunc) HandlerFunc {
	return func(c Context) {
		fn(c.Response(), c.Request())
	}
}

func (c *defaultContext) Request() *http.Request {
	return c.r
}

func (c *defaultContext) Response() http.ResponseWriter {
	return c.w
}

func (c *defaultContext) Context() context.Context {
	return c.r.Context()
}

func (c *defaultContext) Query(query string) string {
	return c.Request().PathValue(query)
}

func (c *defaultContext) XML(status int, data any) error {
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

func (c *defaultContext) JSON(status int, data any) error {
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

func (c *defaultContext) File(path string) {
	http.ServeFile(c.Response(), c.Request(), path)
}

func (c *defaultContext) String(status int, data string) {
	b := bytes.NewBuffer([]byte(data))
	c.w.Header().Set("Content-Type", "text/plain")
	c.w.WriteHeader(status)
	_, _ = b.WriteTo(c.w)
}

func (c *defaultContext) HTML(status int, data string) {
	b := bytes.NewBuffer([]byte(data))
	c.w.Header().Set("Content-Type", "text/html")
	c.w.WriteHeader(status)
	_, _ = b.WriteTo(c.w)
}
