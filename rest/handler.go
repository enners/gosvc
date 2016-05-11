// (c) Knut Enners, see enclosed license

package main

import (
	"fmt"
	"net/http"
)

// Says 'Hello' to last path element
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello, %s.", r.URL.Path[1:])
	} else {
		w.Header().Set("Allowed", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Says 'Hola' to last path element
func handler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hola, %s.", r.URL.Path[1:])
	} else {
		w.Header().Set("Allowed", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
