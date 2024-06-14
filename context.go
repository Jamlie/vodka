package vodka

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Context interface {
	Request() *http.Request
	Response() http.ResponseWriter
	Context() context.Context
	Query(string) string
	XML(int, any) error
	JSON(int, any) error
	File(string)
	String(int, string) error
	HTML(int, string) error
	Url() *url.URL
	ParseForm() error
	FormValue(string) string
	ParseFile(string) (multipart.File, *multipart.FileHeader, error)
	Body() io.ReadCloser
	Method() string
}

type HandlerFunc func(c Context) error

type ctx struct {
	w      http.ResponseWriter
	r      *http.Request
	url    *url.URL
	method string
	body   io.ReadCloser
}

func Wrap(fn http.HandlerFunc) HandlerFunc {
	return func(c Context) error {
		fn(c.Response(), c.Request())
		return nil
	}
}

func (c *ctx) Request() *http.Request {
	return c.r
}

func (c *ctx) Response() http.ResponseWriter {
	return c.w
}

func (c *ctx) Context() context.Context {
	return c.r.Context()
}

func (c *ctx) Query(query string) string {
	return c.Request().PathValue(query)
}

func (c *ctx) XML(status int, data any) error {
	w := bytes.NewBuffer(make([]byte, 0))
	err := xml.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}
	c.w.Header().Set("Content-Type", "application/xml")
	c.w.WriteHeader(status)
	_, err = w.WriteTo(c.w)
	return err
}

func (c *ctx) JSON(status int, data any) error {
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

func (c *ctx) File(path string) {
	http.ServeFile(c.Response(), c.Request(), path)
}

func (c *ctx) String(status int, data string) error {
	b := bytes.NewBuffer([]byte(data))
	c.w.Header().Set("Content-Type", "text/plain")
	c.w.WriteHeader(status)
	_, err := b.WriteTo(c.w)
	return err
}

func (c *ctx) HTML(status int, data string) error {
	b := bytes.NewBuffer([]byte(data))
	c.w.Header().Set("Content-Type", "text/html")
	c.w.WriteHeader(status)
	_, err := b.WriteTo(c.w)
	return err
}

func (c *ctx) Url() *url.URL {
	return c.url
}

func (c *ctx) ParseForm() error {
	return c.Request().ParseForm()
}

func (c *ctx) FormValue(key string) string {
	return c.Request().FormValue(key)
}

func (c *ctx) ParseFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return c.Request().FormFile(key)
}

func (c *ctx) Body() io.ReadCloser {
	return c.body
}

func (c *ctx) Method() string {
	return c.method
}
