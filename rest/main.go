// Copyright information goes here

// first restful service
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	flagHTTP = flag.String("http", ":8080", "HTTP listen address, e.g. localhost:6060")
)

func main() {
	flag.Parse()

	// 12factor: use env var if present, then flag
	httpAddr := ""
	if envHTTP := os.Getenv("HELLO_HTTP"); envHTTP != "" && *flagHTTP == ":8080" {
		httpAddr = envHTTP
	} else {
		httpAddr = *flagHTTP
	}
	http.HandleFunc("/hola/", handler2)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
