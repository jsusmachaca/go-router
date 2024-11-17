package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	ServeMux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		ServeMux: http.NewServeMux(),
	}
}

func (router *Router) Get(path string, handler http.Handler, middleware ...func(next http.Handler) http.Handler) {
	pattern := fmt.Sprintf("GET %s", path)

	if len(middleware) == 0 {
		router.ServeMux.Handle(pattern, handler)
		return
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	router.ServeMux.Handle(pattern, handler)
}

func (router *Router) Post(path string, handler http.Handler, middleware ...func(next http.Handler) http.Handler) {
	pattern := fmt.Sprintf("POST %s", path)

	if len(middleware) == 0 {
		router.ServeMux.Handle(pattern, handler)
		return
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	router.ServeMux.Handle(pattern, handler)
}

func (router *Router) Put(path string, handler http.Handler, middleware ...func(next http.Handler) http.Handler) {
	pattern := fmt.Sprintf("PUT %s", path)

	if len(middleware) == 0 {
		router.ServeMux.Handle(pattern, handler)
		return
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	router.ServeMux.Handle(pattern, handler)
}

func (router *Router) Patch(path string, handler http.Handler, middleware ...func(next http.Handler) http.Handler) {
	pattern := fmt.Sprintf("PATCH %s", path)

	if len(middleware) == 0 {
		router.ServeMux.Handle(pattern, handler)
		return
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	router.ServeMux.Handle(pattern, handler)
}

func (router *Router) Delete(path string, handler http.Handler, middleware ...func(next http.Handler) http.Handler) {
	pattern := fmt.Sprintf("DELETE %s", path)

	if len(middleware) == 0 {
		router.ServeMux.Handle(pattern, handler)
		return
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	router.ServeMux.Handle(pattern, handler)
}

func (router *Router) Options(path string, handler http.Handler, middleware ...func(next http.Handler) http.Handler) {
	pattern := fmt.Sprintf("OPTIONS %s", path)

	if len(middleware) == 0 {
		router.ServeMux.Handle(pattern, handler)
		return
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	router.ServeMux.Handle(pattern, handler)
}

func (router *Router) Connect(path string, handler http.Handler, middleware ...func(next http.Handler) http.Handler) {
	pattern := fmt.Sprintf("CONNECT %s", path)

	if len(middleware) == 0 {
		router.ServeMux.Handle(pattern, handler)
		return
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	router.ServeMux.Handle(pattern, handler)
}

func (router *Router) Trace(path string, handler http.Handler, middleware ...func(next http.Handler) http.Handler) {
	pattern := fmt.Sprintf("TRACE %s", path)

	if len(middleware) == 0 {
		router.ServeMux.Handle(pattern, handler)
		return
	}

	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}

	router.ServeMux.Handle(pattern, handler)
}
