package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "test page")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello, world!\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, header := range req.Header {
		fmt.Fprintln(w, name, header)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":12345", nil)
}