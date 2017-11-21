package service

import (
	"net/http"
	"os"
	"webServer2/service/router"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		Extensions: []string{".html"},
		IndentJSON: true,
	})
	//获取默认的经典的中间件
	n := negroni.Classic()
	//新建一个路由
	mx := mux.NewRouter()

	inition(n, mx, formatter)

	n.UseHandler(mx)
	return n
}

func inition(n *negroni.Negroni, mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		//Getwd是获取当前的目录的字符串路径
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}

	static := negroni.NewStatic(http.Dir(webRoot + "/static"))
	static.Prefix = "/static"
	n.Use(static)
	mx.HandleFunc("/dynamic", router.DynamicHandler)
	mx.HandleFunc("/register", registerHandler(formatter)).Methods("GET")
	mx.HandleFunc("/information", informationHandler(formatter)).Methods("POST")
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(webRoot+"/dynamic/"))))
	mx.PathPrefix("/").HandlerFunc(router.NotFound)
}
