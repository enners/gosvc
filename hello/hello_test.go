package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	var tests = []struct {
		url    string
		method string
		code   int
		body   string
	}{
		{"http://localhost:8080/", "GET", http.StatusOK, "Hello, ."},
		{"http://localhost:8080/World", "GET", http.StatusOK, "Hello, World."},
		{"http://localhost:8080/World", "POST", http.StatusMethodNotAllowed, ""},
	}

	for _, test := range tests {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest(test.method, test.url, nil)
		handler(rr, req)
		if rr.Code != test.code {
			t.Errorf("HTTP %s %s -  code = %d, want %d", test.method, test.url, rr.Code, test.code)
		}
	}
}
