package main

import (
	"fmt"
	"net/http"
)

func queryHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	
	id := values.Get("id")
	pw := values.Get("password")
	
	fmt.Fprintf(w, "ID is %s, password is %s", id, pw)
}

func main() {
	
	myMux := http.NewServeMux()
	myMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Index Page!")
	})
	myMux.HandleFunc("/test", queryHandler)
	http.ListenAndServe(":12345", myMux)
}