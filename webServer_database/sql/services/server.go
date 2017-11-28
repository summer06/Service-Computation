package service

import (
	"net/http"

	"webServer_database/sql/entities"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)
	entities.Initial()

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
	mx.HandleFunc("/service/userinfo", postUserInfoHandler(formatter)).Methods("POST")
	mx.HandleFunc("/service/userinfo", getUserInfoHandler(formatter)).Methods("GET")
	mx.HandleFunc("/service/deleteByID", deleteUserInfoHandler(formatter)).Methods("POST")
	mx.HandleFunc("/service/deleteAll", deleteAllUserInfoHandler(formatter)).Methods("GET")
	mx.HandleFunc("/service/modifyByID", modifyuserinfoHandler(formatter)).Methods("POST")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
	}
}
