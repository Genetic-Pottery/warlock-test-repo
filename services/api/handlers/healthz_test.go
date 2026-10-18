package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Healthz must satisfy http.HandlerFunc.
var _ http.HandlerFunc = Healthz

func TestHealthzReturnsOK(t *testing.T) {
	rec := httptest.NewRecorder()
	Healthz(rec, httptest.NewRequest("GET", "/healthz", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
	if body := rec.Body.String(); body != "ok" {
		t.Fatalf("expected body %q, got %q", "ok", body)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "text/plain; charset=utf-8" {
		t.Fatalf("expected Content-Type %q, got %q", "text/plain; charset=utf-8", ct)
	}
}
