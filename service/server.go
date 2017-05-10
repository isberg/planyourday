package service

import (
	"github.com/urfave/negroni"
	"github.com/gorilla/mux" 
	"github.com/unrolled/render" 
)

func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/projects/{project}", getProjectHandler(formatter)).Methods("GET")
	mx.HandleFunc("/projects", projectCollectionHandler(formatter)).Methods("GET")
	mx.HandleFunc("/projects", createProjectHandler(formatter)).Methods("POST")
}
