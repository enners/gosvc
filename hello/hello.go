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
	port string
)

func init() {
	const (
		portShort   = "p"
		portLong    = "port"
		portDefault = "8080"
		portUsage   = "HTTP listen port"
	)
	flag.Usage = func() {
		usage := fmt.Sprintf("Usage: %s [OPTION]\nStarts simple hello service.\n", os.Args[0])
		usage += "\nOptions are:\n\n"
		usage += "  -" + portShort + ", --" + portLong + "\tHTTP port\n"
		usage += "  -h, --help\tShow this help message\n\n"
		fmt.Fprintf(os.Stderr, usage)
	}
	flag.StringVar(&port, portShort, portDefault, portUsage)
	flag.StringVar(&port, portLong, portDefault, portUsage)
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
