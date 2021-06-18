package cors

import (
	"github.com/michaeliyke/Golang/log"
	"net/http"
	"time"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		log.Println("before handler: middleware starts")
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Printf("MiddleWare finished, %s \n", time.Since(start))
	})
}
