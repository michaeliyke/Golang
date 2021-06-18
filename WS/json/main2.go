package main

import (
	"encoding/json"

	"github.com/michaeliyke/Golang/log"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	PricePerUnit   string `json:"pricePerUnit"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func main() {
	product := &Product{
		ProductID:      123,
		Manufacturer:   "Big Box Comapany",
		PricePerUnit:   "12.99",
		Sku:            "4561qHJK",
		Upc:            "414654444566",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(productJSON))
}
