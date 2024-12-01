package api

import "net/http"

func registerRoutes(h *http.ServeMux) {
	h.HandleFunc("GET /ping", ping)

	h.HandleFunc("POST /account/make-session", makeSession)
	h.HandleFunc("POST /account/create", createAccount)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}

func makeSession(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented"))
}
