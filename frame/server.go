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
	s.Use(Logger(), Cors())
	return s
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *Server) Use(middlewares ...HandlerFunc) {
	s.middlewares = append(s.middlewares, middlewares...)
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
		// 404
		var mergedHandlers HandlersChain
		if len(s.middlewares) > 0 {
			finalSize := len(s.middlewares) + len(NotFoundHandler)
			mergedHandlers = make(HandlersChain, finalSize)
			copy(mergedHandlers, s.middlewares)
			copy(mergedHandlers[len(s.middlewares):], NotFoundHandler)
		} else {
			// middleware empty
			finalSize := len(NotFoundHandler)
			mergedHandlers = make(HandlersChain, finalSize)
			copy(mergedHandlers, NotFoundHandler)
		}
		c.handlers = mergedHandlers
	}
	c.Next()
}

var NotFoundHandler = []HandlerFunc{func(c *Context) {
	c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
}}

func (s *Server) POST(pattern string, handler HandlerFunc) {
	s.router.addRoute(http.MethodPost, pattern, []HandlerFunc{handler})
}

func (s *Server) DELETE(pattern string, handler HandlerFunc) {
	s.router.addRoute(http.MethodDelete, pattern, []HandlerFunc{handler})
}

func (s *Server) GET(pattern string, handler HandlerFunc) {
	s.router.addRoute(http.MethodGet, pattern, []HandlerFunc{handler})
}
func (s *Server) PUT(pattern string, handler HandlerFunc) {
	s.router.addRoute(http.MethodPut, pattern, []HandlerFunc{handler})
}

func (s *Server) PATCH(pattern string, handler HandlerFunc) {
	s.router.addRoute(http.MethodPatch, pattern, []HandlerFunc{handler})
}
