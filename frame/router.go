package frame

type HandlerFunc func(*Context)

type router struct {
	middlewareChain []HandlerFunc
	handlers        map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// Use user middleware
func (r *router) Use(middleware ...HandlerFunc) {
	r.middlewareChain = append(r.middlewareChain, middleware...)
}
