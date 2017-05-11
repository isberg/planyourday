package service

import (
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

func getProjectHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
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

type project struct {
	Name string
}

func createProjectHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		println("payload: " + string(payload))

		var project project
		var name string
		name = "noname"
		if err:=json.Unmarshal(payload, &project); err== nil {
			name = project.Name
		} else {
			println("error:", err)
		}

		println("name: ", name)
		w.Header().Set("Location", "/projects/" + name)
		formatter.JSON(w, http.StatusCreated, nil)
	}
}
