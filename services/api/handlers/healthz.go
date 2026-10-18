package handlers

import "net/http"

// Healthz reports that the process is up. It answers 200 with the
// plain-text body "ok" and checks no dependencies; the Router decides
// which methods reach it.
func Healthz(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
