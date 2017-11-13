package mux

import (
	"io"
	"net/http"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

type MyHandler struct{}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = hello
	h, ok := mux[r.URL.String()]
	if ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My Server: "+r.URL.String())
}
