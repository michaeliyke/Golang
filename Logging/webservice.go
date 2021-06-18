package main

import (
	"encoding/json"
	"net/http"

	"github.com/michaeliyke/Golang/log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		m := map[string]string{"name": name}
		enc := json.NewEncoder(w)
		enc.Encode(m)
	})

	error := http.ListenAndServe(":80", nil)
	if error != nil {
		log.Fatal("ERROR OCCURED", error)
	}
}
