package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestHandler(t *testing.T) {
	expected := "Yo Yo Yo."

	req := httptest.NewRequest(http.MethodGet, "/api", nil)
	w := httptest.NewRecorder()

	RequestHandler(w, req)
	res := w.Result()
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Expected %v but got %v", expected, err)
	}
	if err != nil || string(body) != expected {
		t.Errorf("Expected %v but got %v", expected, string(body))
	}
}
