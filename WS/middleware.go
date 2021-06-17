package main

import (
	"log"
	"net/http"
	"time"
)

func middleWareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("before handler: middleware starts")
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Println("MiddleWare finished, %s", time.Since(start))
	})
}
