package router

import (
  "io"
  "net/http"
  "html/template"
  "fmt"
  // "github.com/gorilla/mux"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
  panic(fmt.Sprintf(
		"Error: unknown path: %s, not implemented", r.URL.Path))
}

func DynamicHandler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("dynamic/index.html")
  if (err != nil) {
    io.WriteString(w, "page fault")
  } else {
    t.Execute(w, nil)
  }
}
