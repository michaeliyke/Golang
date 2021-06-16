package main

import (
	"net/http"
)

type fooHandler struct {
	Message string
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bar called"))
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "Hello world!"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":80", nil)
}
