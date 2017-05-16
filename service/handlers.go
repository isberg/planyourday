package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"io/ioutil"
	"net/http"
)

type project struct {
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	Steps []string `json:"steps"`
}

func getProjectHandler(formatter *render.Render) http.HandlerFunc {
	learndocker := project { Name:"Learn Docker", Steps:([]string{}) }

	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		println("project:", project)
		if project == "learndocker" {
			formatter.JSON(w, http.StatusOK, learndocker)
		} else {
			formatter.Text(w, http.StatusNotFound, "")
		}
	}
}

func projectCollectionHandler(formatter *render.Render) http.HandlerFunc {
	projects := []project {project {Name:"Learn Go", Steps:[]string{}}, project{Name:"Learn Docker", Steps:[]string{}} }
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, projects)
	}
}

func createProjectHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, _ := ioutil.ReadAll(req.Body)
		println("payload: " + string(payload))

		var newproject project
		var name string
		name = "noname"
		if err := json.Unmarshal(payload, &newproject); err == nil {
			name = newproject.Name
		} else {
			println("error:", err)
		}

		println("name: ", name)
		w.Header().Set("Location", "/projects/"+name)

		data, _ := json.Marshal(newproject)
		println("data:", string(data))
		formatter.JSON(w, http.StatusCreated, newproject)
	}
}
