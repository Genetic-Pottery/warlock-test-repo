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
			{method: "GET", path: "/invoices/{id}"},
		},
	}
}

// Handle reports whether method and path match a registered route.
// The router's prefix is stripped from path before matching, since
// routes are stored prefix-free; a path without the prefix matches
// nothing. The method comparison is exact and case-sensitive. Paths
// match segment by segment: see matchPath.
func (r *Router) Handle(method, path string) bool {
	rest, ok := strings.CutPrefix(path, r.prefix)
	if !ok {
		return false
	}
	for _, candidate := range r.routes {
		if candidate.method == method && matchPath(candidate.path, rest) {
			return true
		}
	}
	return false
}

// matchPath reports whether path matches pattern segment by segment.
// A pattern segment of the form "{name}" matches any one non-empty
// segment; every other segment must match exactly and case-sensitively.
func matchPath(pattern, path string) bool {
	want := strings.Split(pattern, "/")
	got := strings.Split(path, "/")
	if len(want) != len(got) {
		return false
	}
	for i, seg := range want {
		if strings.HasPrefix(seg, "{") && strings.HasSuffix(seg, "}") {
			if got[i] == "" {
				return false
			}
			continue
		}
		if seg != got[i] {
			return false
		}
	}
	return true
}
