package server

import (
	"net/http"
	"webServer/server/mux"
)

func Serve() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &mux.MyHandler{},
	}
	server.ListenAndServe()
}
