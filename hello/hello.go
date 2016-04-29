// Copyright information goes here

// Simple hello service to cover basic go stuff
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	flagHTTP = flag.String("http", ":8080", "HTTP listen address, e.g. localhost:6060")
)

// Says 'Hello' to last path element
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s.", r.URL.Path[1:])
}

func main() {
	flag.Parse()

	// 12factor: use env var if present, then flag
	httpAddr := ""
	if envHTTP := os.Getenv("HELLO_HTTP"); envHTTP != "" && *flagHTTP == ":8080" {
		httpAddr = envHTTP
	} else {
		httpAddr = *flagHTTP
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
