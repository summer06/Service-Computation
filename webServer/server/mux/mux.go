package mux

import (
	"io"
	"net/http"
)

//use a map to record the information of routers
var mux map[string]func(http.ResponseWriter, *http.Request)

//a struct that implement ServeHTTP function int the interface
type MyHandler struct{}

//main page handler
func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

//handler of /name
func name(w http.ResponseWriter, r *http.Request) {
	//get the query string
	myname := r.URL.Query().Get("myname")
	if myname == "" {
		io.WriteString(w, "Your name value is empty")
	} else {
		io.WriteString(w, "Your name is "+myname)
	}
}

func (*MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	//add router information and its handler
	mux["/"] = hello
	mux["/name"] = name
	//return handler according to the path
	h, ok := mux[r.URL.Path]
	if ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My Server: "+r.URL.String())
}
