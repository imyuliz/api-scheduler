package frame

import (
	"net/http"
)

type Server struct {
	middlewares HandlersChain // support middleware
	router      *router
}

func NewServer() *Server {
	s := &Server{router: newRouter()}
	return s
}

func NewDefaultServer() *Server {
	s := NewServer()
	s.Use(Logger())
	return s
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	s.request(c)

}

func (s *Server) request(c *Context) {
	node, _ := s.router.getRoute(c.Method, c.Path)
	if node != nil {
		key := c.Method + "-" + node.pattern
		c.handlers = s.router.handlers[key]
	} else {
		var mergedHandlers HandlersChain
		if len(s.middlewares) > 0 {
			finalSize := len(s.middlewares) + 1
			mergedHandlers = make(HandlersChain, finalSize)
			copy(mergedHandlers, s.middlewares)
			mergedHandlers = append(mergedHandlers, func(c *Context) {
				c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
			})
		} else {
			mergedHandlers = make(HandlersChain, 0)
			mergedHandlers = append(mergedHandlers, func(c *Context) {
				c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
			})
		}
		c.handlers = mergedHandlers
	}
	c.Next()
}

func (s *Server) GET(pattern string, handler HandlerFunc) {
	s.router.addRoute("GET", pattern, []HandlerFunc{handler})
}

func (s *Server) Use(middlewares ...HandlerFunc) {
	s.middlewares = append(s.middlewares, middlewares...)
}
