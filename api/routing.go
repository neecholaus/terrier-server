package api

import "net/http"

func registerRoutes(h *http.ServeMux) {
	h.HandleFunc("GET /ping", ping)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong\n"))
}
