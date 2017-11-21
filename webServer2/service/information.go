package service

import (
	"net/http"

	"github.com/unrolled/render"
)

func informationHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
    req.ParseForm()
		formatter.HTML(w, http.StatusOK, "information", struct {
			Name      string `json:"name"`
			Id        string `json:"id"`
      Sex       string `json:"sex"`
      Class     string `json:"class"`
		}{
      Name: req.Form["name"][0],
      Id: req.Form["id"][0],
      Sex: req.Form["sex"][0],
      Class: req.Form["class"][0],
    })
	}
}
