package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var serviceYUrl = "http://localhost:8081"
var useServiceYContractServer = true

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request")
		return
	}
	name := query.Get("name")
	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "You must supply a name")
		return
	}

	serviceYResponse, err := CallServiceY()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Bad request")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s. %s", name, serviceYResponse)
}

func CallServiceY() (string, error) {
	url := fmt.Sprintf("%s/api", serviceYUrl)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func main() {
	if useServiceYContractServer {
		server := StartServiceY()
		defer server.Close()
		serviceYUrl = server.URL
	}

	http.HandleFunc("/api", RequestHandler)
	http.ListenAndServe(":8080", nil)
}
