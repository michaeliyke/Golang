package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func productHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJSON, error := json.Marshal(productList)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "aplication/json")
		w.Write(productsJSON)
	case http.MethodPost:
		// Create a bew product
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		newProduct.ProductID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		log.Println(newProduct)
		return
	}
}

func main() {
	log.Println("Hello")
	http.HandleFunc("/products", productHandler)
	http.ListenAndServe(":80", nil)
}
