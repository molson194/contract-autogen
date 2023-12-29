package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	expected := "Hello john. Yo Yo Yo."

	server := StartServiceY()
	defer server.Close()

	serviceYUrl = server.URL

	req := httptest.NewRequest(http.MethodGet, "/api?name=john", nil)
	w := httptest.NewRecorder()

	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Expected %v but got %v", expected, err)
		return
	}
	if string(body) != expected {
		t.Errorf("Expected %v but got %v", expected, string(body))
	}
}
