package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProjectReturns200ForExistingProject(t *testing.T) {
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	server := NewServer()
	targetProject := "learngo"

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/projects/"+targetProject, nil)
	server.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}
}

func TestGetProjectReturns404ForNonExistingProject(t *testing.T) {
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	server := NewServer()
	targetProject := "nothing"

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/projects/"+targetProject, nil)
	server.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusNotFound {
		t.Errorf("Expected %v; received %v", http.StatusNotFound, recorder.Code)
	}
}

func TestListAllProjectsReturns200(t *testing.T) {
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	server := NewServer()

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/projects", nil)
	server.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}
}

func TestCreateProjectReturns201ForNewProject(t *testing.T) {
	var request *http.Request
	var recorder *httptest.ResponseRecorder

	server := NewServer()

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("POST", "/projects", nil)
	server.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected %v; received %v", http.StatusCreated, recorder.Code)
	}

	location := recorder.Header().Get("Location")
	println("location:" + location)
	if location == "" {
		t.Errorf("Expected Location Header to be set")
	}
}
