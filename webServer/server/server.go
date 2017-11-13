package server

import (
	"net/http"
	"webServer/server/mux"
)

func Serve(port string) {
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux.Mux(),
	}
	server.ListenAndServe()
}
