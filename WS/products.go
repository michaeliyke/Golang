package main

import (
	"encoding/json"

	"example.com/log"
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

var productList []Product

func init() {
	productsJSON := `[
		{
			"productId": 1,
			"manufacturer": "Johns-Jenkins",
			"sku": "p5z343vdS",
			"upc": "939581000000",
			"pricePerUnit": "497.45",
			"quantityOnHand": 9703,
			"productName": "sticky note"
		},
		{
			"productId": 2,
			"manufacturer": "Hassel, Schimel and Feeney",
			"sku": "i7v300kmx",
			"upc": "740979000000",
			"pricePerUnit": "282.29",
			"quantityOnHand": 9217,
			"productName": "leg warmers"
		},
		{
			"productId": 3,
			"manufacturer": "Swaniawski, Bartoletti and Bruen",
			"sku": "q0L657ys7",
			"upc": "111730000000",
			"pricePerUnit": "436.26",
			"quantityOnHand": 5905,
			"productName": "Lamp shade"
		},
		{
			"productId": 4,
			"manufacturer": "Big Box Comapany",
			"pricePerUnit": "12.99",
			"sku": "4561qHJK",
			"upc": "414654444566",
			"quantityOnHand": 28,
			"productName": "Gizmo"
		},
		{
        "productId": 5,
        "manufacturer": "Small Box Comapany",
        "pricePerUnit": "12.99",
        "sku": "4hslj90JKL",
        "upc": "424654445566",
        "quantityOnHand": 18,
        "productName": "Sprocket"
    }
	]`

	error := json.Unmarshal([]byte(productsJSON), &productList)
	if error != nil {
		log.Fatal(error)
	}
}

func getNextID() int {
	highestID := -1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}
	return highestID + 1
}

func findPoductByID(productID int) (*Product, int) {
	for i, product := range productList {
		if product.ProductID == productID {
			return &product, i
		}
	}
	return nil, 0
}
