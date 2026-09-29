package handlers

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

func (r *Router) Handle(path string) bool {
	return len(path) > 0
}
