package gee

type Router struct {
	handlers map[string]HandlerFunc
}

func (router *Router) addRoute(method, pattern string, handler HandlerFunc)  {
	k := method + "-" + pattern
	router.handlers[k] = handler
}

func (router *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler,ok := router.handlers[key];ok {
		handler(c)
	} else {

	}
}

func NewRouter () *Router {
	return &Router{handlers:make(map[string]HandlerFunc)}
}

