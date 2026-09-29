package handlers

import "strings"

// A single registered route: an HTTP method paired with a
// prefix-free path, so DefaultPrefix can change independently.
type route struct {
	method string
	path   string
}

// A registry of HTTP routes.
type Router struct {
	prefix string
	routes []route
}

const DefaultPrefix = "/v1"

func NewRouter() *Router {
	return &Router{
		prefix: DefaultPrefix,
		routes: []route{
			{method: "POST", path: "/entries/reverse"},
		},
	}
}

// Handle reports whether method and path match a registered route.
// The router's prefix is stripped from path before matching, since
// routes are stored prefix-free; a path without the prefix matches
// nothing. Both comparisons are exact and case-sensitive.
func (r *Router) Handle(method, path string) bool {
	rest, ok := strings.CutPrefix(path, r.prefix)
	if !ok {
		return false
	}
	for _, candidate := range r.routes {
		if candidate.method == method && candidate.path == rest {
			return true
		}
	}
	return false
}
