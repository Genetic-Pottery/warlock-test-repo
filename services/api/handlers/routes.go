package handlers

// A single HTTP route, identified by its method and path.
type Route struct {
	Method string
	Path   string
}

// A registry of HTTP routes.
type Router struct {
	prefix string
	routes []Route
}

const DefaultPrefix = "/v1"

func NewRouter() *Router {
	r := &Router{prefix: DefaultPrefix}
	r.routes = []Route{
		{Method: "GET", Path: r.prefix + "/health"},
		{Method: "GET", Path: r.prefix + "/things"},
		{Method: "POST", Path: r.prefix + "/things"},
	}
	return r
}

func (r *Router) Handle(path string) bool {
	return len(path) > 0
}
