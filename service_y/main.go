package main

import (
	"fmt"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Yo Yo Yo.")
}

func main() {
	http.HandleFunc("/api", RequestHandler)
	http.ListenAndServe(":8081", nil)
}
