package handlers

import "testing"

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
