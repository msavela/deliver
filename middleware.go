package deliver

import (
	"net/http"
)

type Middleware struct {
	handler HandlerNext
}

type (
	HandlerNext interface {
		ServeHTTP(Response, *Request, func())
	}

	HttpHandlerNext interface {
		ServeHTTP(http.ResponseWriter, *http.Request, func())
	}
)

type MiddlewareHandlerFunc func(res Response, req *Request, next func())
func (h MiddlewareHandlerFunc) ServeHTTP(res Response, req *Request, next func()) {
	h(res, req, next)
}

func NewMiddleware(handler HandlerNext) *Middleware {
	return &Middleware{handler}
}

func NewMiddlewareHandler(handler http.Handler) *Middleware {
	m := &Middleware{}
	m.handler = MiddlewareHandlerFunc(func(res Response, req *Request, next func()) {
		handler.ServeHTTP(res.ResponseWriter, req.Request)
		next()
	})
	return m
}

func NewMiddlewareHandlerNext(handler HttpHandlerNext) *Middleware {
	m := &Middleware{}
	m.handler = MiddlewareHandlerFunc(func(res Response, req *Request, next func()) {
		handler.ServeHTTP(res.ResponseWriter, req.Request, next)
	})
	return m
}