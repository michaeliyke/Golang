package main

import (
	"log"
	"net/http"
	"github.com/michaeliyke/Golang/webservice/product"
)

const apiBasePath = "/api"

func main() {
	log.Println("Hello")
	product.setUpRoutes(apiBasePath)
	http.ListenAndServe(":80", nil)
}
