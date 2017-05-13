package service

import (
	"strings"
	"testing"

	"encoding/json"
	"net/http"
)

func TestCreateProjectReturns201ForNewProject(t *testing.T) {
	recorder := makeRequest("POST", "/projects", "{}")

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected %v; received %v", http.StatusCreated, recorder.Code)
	}
}

func TestCreateProjectSetsLocationHeader(t *testing.T) {
	recorder := makeRequest("POST", "/projects", "{}")

	location := recorder.Header().Get("Location")
	if location == "" {
		t.Errorf("Expected Location Header to be set")
	}
}

func TestCreateProjectSetsLocationHeaderThatMatchesProjectName(t *testing.T) {
	name := "myproject"

	json := "{ \"name\": \"" + name + "\" }"
	println("json:" + json)

	recorder := makeRequest("POST", "/projects", json)

	location := recorder.Header().Get("Location")
	if !strings.HasSuffix(location, name) {
		t.Errorf("Expected Location Header '%v' to match project name '%v'", location, name)
	}
}

func TestCreateProjectReturnsTheNewProject(t *testing.T) {
	myproject := project {"myproject"} 
	data, _ := json.Marshal(myproject) 
	recorder := makeRequest("POST", "/projects", string(data))

	var newProject project
	if err := json.Unmarshal(recorder.Body.Bytes(), &newProject); err != nil {
		println("error:", err.Error())
		body := recorder.Body.String()
		t.Error("Did not receive expected response body, received", body)
		return
	}

	if (myproject.Name != newProject.Name) {
		t.Error("Mismatch in name, expected:", myproject.Name, "but was", newProject.Name)
	}
}
