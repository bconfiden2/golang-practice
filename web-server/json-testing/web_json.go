package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Student struct {
	Name string
	Age int
	Score int
}

func main() {
	http.ListenAndServe(":12345", MakeHandler())
}

func MakeHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentDataHandler)
	return mux
}

func StudentDataHandler(w http.ResponseWriter, r *http.Request) {
	student := Student{"bconfiden2", 27, 100}
	data, _ := json.Marshal(student)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}