package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

func StartServiceY() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api" {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Yo Yo Yo.")
		}
	}))
	return server
}
