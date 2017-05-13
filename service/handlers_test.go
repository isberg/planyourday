package service

import (
	"net/http"
	"testing"
)

func TestGetProjectReturns200ForExistingProject(t *testing.T) {
	recorder := makeRequest("GET", "/projects/learndocker", "")

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected Status Code %v, but was %v", http.StatusOK, recorder.Code)
	}
}

func TestGetProjectReturns404ForNonExistingProject(t *testing.T) {
	recorder := makeRequest("GET", "/projects/nothing", "")

	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected %v; received %v", http.StatusNotFound, recorder.Code)
	}

	if body := recorder.Body.String(); len(body) > 0 {
		t.Errorf("Expected no body; received '%v'", body)
	}
}

func TestListAllProjectsReturns200(t *testing.T) {
	recorder := makeRequest("GET", "/projects", "")

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}
}
