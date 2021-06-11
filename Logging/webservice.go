package main

import (
	"net/http"

	"example.com/log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	error := http.ListenAndServe(":3000", nil)
	if error != nil {
		log.Fatal("ERROR OCCURED", error)
	}
}
