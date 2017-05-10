package service

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func getProjectHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		println("project: " + project)
		if (project == "learngo") {
			formatter.JSON(w, http.StatusOK, nil)
		} else {
			formatter.JSON(w, http.StatusNotFound, nil)
		}
	}
}
