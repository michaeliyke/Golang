package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Hello")
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productHandler)
	http.ListenAndServe(":80", nil)
}
