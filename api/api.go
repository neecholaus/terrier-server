package api

import (
	"fmt"
	"net/http"
	"nick/terrier-server/system"
)

func Serve() {
	port := 80
	system.Logger.Info(fmt.Sprintf("starting API on port %d", port))

	handler := http.NewServeMux()

	s := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	registerRoutes(handler)

	s.ListenAndServe()
}
