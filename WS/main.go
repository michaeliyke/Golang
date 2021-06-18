package main

import (
	"github.com/michaeliyke/Golang/log"
	"net/http"

	"github.com/michaeliyke/Golang/WS/product"
)

const apiBasePath = "/api"

func main() {
	log.Println("Hello")
	product.SetUpRoutes(apiBasePath)
	http.ListenAndServe(":80", nil)
}
