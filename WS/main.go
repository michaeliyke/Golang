package main

import (
	"log"
	"net/http"

	"github.com/michaeliyke/Golang/WS/product"
)

const apiBasePath = "/api"

func main() {
	log.Println("Hello")
	product.SetUpRoutes(apiBasePath)
	http.ListenAndServe(":80", nil)
}
