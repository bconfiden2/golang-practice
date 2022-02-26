package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "HTTPS TEST")
	})
	if err := http.ListenAndServeTLS(":12345", "server.crt", "server.key", nil) ; err != nil {
		log.Fatal(err)
	}
}