package service

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

func getProjectHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		println("project: " + project)
		if project == "learngo" {
			formatter.JSON(w, http.StatusOK, nil)
		} else {
			formatter.JSON(w, http.StatusNotFound, nil)
		}
	}
}

func projectCollectionHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, nil)
	}
}

func createProjectHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Location", "Here")
		formatter.JSON(w, http.StatusCreated, nil)
	}
}
