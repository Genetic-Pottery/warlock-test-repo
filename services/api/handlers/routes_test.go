package handlers

import "testing"

func TestHandleMatchesRegisteredMethodAndPath(t *testing.T) {
	r := NewRouter()
	if !r.Handle("POST", "/v1/entries/reverse") {
		t.Fatal("expected POST /v1/entries/reverse to be handled")
	}
}

func TestHandleRejectsWrongMethodOnRegisteredPath(t *testing.T) {
	r := NewRouter()
	if r.Handle("GET", "/v1/entries/reverse") {
		t.Fatal("expected GET /v1/entries/reverse to be unhandled: the path is registered, the method is not")
	}
}

func TestHandleRejectsPathWithoutPrefix(t *testing.T) {
	r := NewRouter()
	if r.Handle("POST", "/entries/reverse") {
		t.Fatal("expected POST /entries/reverse to be unhandled: the /v1 prefix is missing")
	}
}

func TestHandleRejectsUnregisteredPath(t *testing.T) {
	r := NewRouter()
	if r.Handle("GET", "/v1/things") {
		t.Fatal("expected GET /v1/things to be unhandled: /things is not a route")
	}
}
