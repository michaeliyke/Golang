package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Hello")
	productListHandler := http.HandlerFunc(productsHandler)
	productItemHandler := http.HandlerFunc(productHandler)
	http.Handle("/products", middleWareHandler(productListHandler))
	http.Handle("/products/", middleWareHandler(productItemHandler))
	http.ListenAndServe(":80", nil)
}
