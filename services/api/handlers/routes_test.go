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

func TestHandleMatchesInvoiceByID(t *testing.T) {
	r := NewRouter()
	if !r.Handle("GET", "/v1/invoices/42") {
		t.Fatal("expected GET /v1/invoices/42 to be handled")
	}
}

func TestHandleRejectsWrongMethodOnInvoice(t *testing.T) {
	r := NewRouter()
	if r.Handle("POST", "/v1/invoices/42") {
		t.Fatal("expected POST /v1/invoices/42 to be unhandled: the path is registered, the method is not")
	}
}

func TestHandleRejectsInvoiceWithoutID(t *testing.T) {
	r := NewRouter()
	for _, path := range []string{"/v1/invoices", "/v1/invoices/"} {
		if r.Handle("GET", path) {
			t.Fatalf("expected GET %s to be unhandled: the {id} segment is missing", path)
		}
	}
}

func TestHandleRejectsInvoiceWithExtraSegment(t *testing.T) {
	r := NewRouter()
	if r.Handle("GET", "/v1/invoices/42/lines") {
		t.Fatal("expected GET /v1/invoices/42/lines to be unhandled: {id} matches one segment")
	}
}
