package geegin

import (
	"fmt"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*Node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*Node),
		handlers: make(map[string]HandlerFunc),
	}
}

// parsePattern
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}

	}
	return parts
}

func (r *router) addRoute(method, path string, handler HandlerFunc) {
	parts := parsePattern(path)
	key := method + "-" + path
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &Node{}
	}
	r.roots[method].insert(path, parts, 0)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.Pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

func (r *router) getRoute(method, path string) (*Node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.Pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				fmt.Println(searchParts[index:])
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
