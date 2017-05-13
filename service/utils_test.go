package service

import (
	"bytes"

	"net/http"
	"net/http/httptest"
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

	body := recorder.Body.String()
	println("recieved body: " + body)

	return recorder
}
