package handlers

// A registry of HTTP routes.
type Router struct {
	prefix string
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
