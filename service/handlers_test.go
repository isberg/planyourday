package service

import (
	"testing"

	"encoding/json"
	"net/http"
)

func TestGetExistingProject(t *testing.T) {
	recorder := makeRequest("GET", "/projects/learndocker", "")

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected Status Code %v, but was %v", http.StatusOK, recorder.Code)
	}

	body := recorder.Body.Bytes()
	var aproject project

	if err := json.Unmarshal(body, &aproject); err != nil {
		t.Errorf("Error '%v' when Unmarshalling '%v'", err.Error, string(body))
		return
	}

	print("body: '" + string(body) + "', name: '"+ aproject.Name + "'")

	if aproject.Name != "Learn Docker" {
		t.Errorf("Project name expected: %v, actual: %v", "Learn Docker", aproject.Name)
	}

	if len(aproject.CreatedAt) != 0 {
		t.Errorf("Project creation time expected actual: %v", aproject.CreatedAt)
	}

	if steps := len(aproject.Steps); steps != 0 {
		t.Errorf("Expected no steps, actual: %v", steps)
	}
}

func TestGetNonExistingProject(t *testing.T) {
	recorder := makeRequest("GET", "/projects/nothing", "")

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected %v; received %v", http.StatusNotFound, recorder.Code)
	}

	if body := recorder.Body.String(); len(body) > 0 {
		t.Errorf("Expected no body; received '%v'", body)
	}
}

func TestListAllProjects(t *testing.T) {
	recorder := makeRequest("GET", "/projects", "")

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}

	var projects []project
	if err:=json.Unmarshal(recorder.Body.Bytes(), &projects); err!=nil {
		t.Fatal(err.Error)
	}

	if body:= recorder.Body.String(); body == "null\n" {
		t.Error("Should return empty json array instead of null")
	}

	if n:=len(projects); n < 2 {
		t.Errorf("Should return 2 default projects in test mode, but returned %v", n)
	}

	if n:=len(projects); n > 0 {
		proj := projects[0]
		if proj.Steps == nil {
			t.Errorf("Steps cannot be null, must be array with zero elements")
		}
	}

}


