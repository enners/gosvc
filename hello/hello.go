// Copyright information goes here

// Simple hello service to cover basic go stuff
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port string
)

func init() {
	const (
		portDefault = "8080"
		portUsage   = "HTTP listen port"
	)
	flag.StringVar(&port, "port", portDefault, portUsage)
	flag.StringVar(&port, "p", portDefault, portUsage)
}

// Says 'Hello' to last path element
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s.", r.URL.Path[1:])
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
