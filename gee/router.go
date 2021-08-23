package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func (router *Router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	_, ok := router.roots[method]
	if !ok {
		router.roots[method] = &node{}
	}
	router.roots[method].insert(pattern, parts, 0)
	k := method + "-" + pattern
	router.handlers[k] = handler
}

func (router *Router) getRoute(method string, pattern string) (*node, map[string]string) {
	parts := parsePattern(pattern)
	root, ok := router.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(parts, 0)
	if nil == n {
		return nil, nil
	}
	params := make(map[string]string)
	for index, part := range parsePattern(n.pattern) {
		if part[0] == ':' {
			params[part[1:]] = parts[index]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(parts[index:], "/")
			break
		}
	}
	return n, params
}

func (router *Router) handle(c *Context) {
	node, params := router.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.pattern
		router.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
}

func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, v := range vs {
		if v != "" {
			parts = append(parts, v)
			if v[0] == '*' {
				break
			}
		}
	}
	return parts
}
