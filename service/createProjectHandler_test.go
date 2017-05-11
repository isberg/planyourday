package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func makeRequest(method string, url string, json string) *httptest.ResponseRecorder {
	var (
		request  *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := NewServer()
	recorder = httptest.NewRecorder()
	if json != "" {
		request, _ = http.NewRequest(method, url, bytes.NewBufferString(json))
	} else {
		request, _ = http.NewRequest(method, url, nil)
	}

	server.ServeHTTP(recorder, request)
	return recorder
}

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
