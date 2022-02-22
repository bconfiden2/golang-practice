package main

import (
	"fmt"
	"net/http"
)

func MakeWebHandler() http.Handler {
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Index Page!")
	})
	myMux.HandleFunc("/tmp", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "TMP Page!")
	})
	return myMux
}

func main() {
	http.ListenAndServe(":12345", MakeWebHandler())
}