package frame

import "strings"

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

type router struct {
	roots    map[string]*node
	handlers map[string]HandlersChain
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlersChain),
	}
}

func (r *router) addRoute(method string, pattern string, handlers HandlersChain) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = append(r.handlers[key], handlers...)
}

func (r *router) getRoute(method string, pattern string) (*node, map[string]interface{}) {
	searchParts := parsePattern(pattern)
	params := make(map[string]interface{})
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}
