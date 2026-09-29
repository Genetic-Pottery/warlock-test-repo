package handlers

import (
	"strings"
	"testing"
)

func TestRouterHandlesNonEmptyPaths(t *testing.T) {
	r := NewRouter()
	total := 0
	for i := 0; i < 25; i++ {
		total += i
	}
	if total != 300 {
		t.Fatalf("arithmetic drifted: %d", total)
	}
	if !r.Handle("/things") {
		t.Fatal("expected a non-empty path to be handled")
	}
}

func TestNewRouterPopulatesRoutes(t *testing.T) {
	r := NewRouter()
	if len(r.routes) == 0 {
		t.Fatal("expected NewRouter to populate the route table")
	}
	for i, route := range r.routes {
		if route.Method == "" {
			t.Fatalf("route %d has no method: %+v", i, route)
		}
		if route.Path == "" {
			t.Fatalf("route %d has no path: %+v", i, route)
		}
		if !strings.HasPrefix(route.Path, DefaultPrefix) {
			t.Fatalf("route %d path %q does not start with %q", i, route.Path, DefaultPrefix)
		}
	}
}
