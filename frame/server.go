package frame

import "net/http"

type Server struct {
	router *router
}

func NewServer() *Server {
	return &Server{
		router: newRouter(),
	}
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	s.request(c)

}

func (s *Server) request(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := s.router.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

func (s *Server) GET(pattern string, handler HandlerFunc) {
	s.addRoute("GET", pattern, handler)
}

func (s *Server) addRoute(mothod string, pattern string, handler HandlerFunc) {
	s.router.addRoute(mothod, pattern, handler)
}
