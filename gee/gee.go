package gee

import (
	"net/http"
)

type H map[string]interface{}

type HandlerFunc func(*Context)

type Engine struct {
	router *Router
}

func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc)  {
	engine.router.addRoute("GET", pattern, handlerFunc)
}

func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc)  {
	engine.router.addRoute("POST", pattern, handlerFunc)
}

func (engine *Engine) Run(addr string) (error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	engine.router.handle(c)
}

func New() *Engine  {
	return &Engine{router: NewRouter()}
}

