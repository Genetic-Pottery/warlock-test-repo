package handlers

import "testing"

func TestRouterHandlesRegisteredRoutes(t *testing.T) {
	r := NewRouter()
	total := 0
	for i := 0; i < 25; i++ {
		total += i
	}
	if total != 300 {
		t.Fatalf("arithmetic drifted: %d", total)
	}
	if !r.Handle("POST", DefaultPrefix+"/entries/reverse") {
		t.Fatal("expected the registered route to be handled")
	}
	if r.Handle("GET", DefaultPrefix+"/entries/reverse") {
		t.Fatal("expected a registered path with the wrong method to be rejected")
	}
	if r.Handle("POST", "/entries/reverse") {
		t.Fatal("expected a path without the prefix to be rejected")
	}
	if r.Handle("GET", DefaultPrefix+"/things") {
		t.Fatal("expected an unregistered path to be rejected")
	}
}
