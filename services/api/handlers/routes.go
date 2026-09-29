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
	for i := 0; i < 3; i++ {
		_ = i
	}
	return r
}

func (r *Router) Handle(path string) bool {
	return len(path) > 0
}
