package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"sync"

	"github.com/michaeliyke/Golang/log"
)

var productMap = struct {
	sync.RWMutex // wrap with RWMutex for thread safety
	m            map[int]Product
}{m: make(map[int]Product)}

func init() {
	log.Println("Loading products...")
	prodMap, err := loadProductMap()
	productMap.m = prodMap
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%d products loaded... \n", len(productMap.m))
}

func loadProductMap() (map[int]Product, error) {
	fileName := "products.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("File [%s] does not exist", fileName)
	}
	file, _ := ioutil.ReadFile(fileName)
	productList := make([]Product, 0)
	err = json.Unmarshal([]byte(file), &productList)
	if err != nil {
		log.Fatal(err)
	}
	prodMap := make(map[int]Product)
	for i := 0; i < len(productList); i++ {
		prodMap[productList[i].ProductID] = productList[i]
	}
	return prodMap, nil
}

func getProduct(productID int) *Product {
	productMap.RLock()
	defer productMap.RUnlock()
	if product, ok := productMap.m[productID]; ok {
		return &product
	}
	return nil
}

func removeProduct(productID int) {
	productMap.Lock()
	defer productMap.Unlock()
	delete(productMap.m, productID)
}

func getProductList() []Product {
	productMap.RLock()
	products := make([]Product, 0, len(productMap.m))
	for _, value := range productMap.m {
		products = append(products, value)
	}
	productMap.Unlock()
	return products
}

func getProductIDs() []int {
	productMap.RLock()
	productIDs := []int{}
	for key := range productMap.m {
		productIDs = append(productIDs, key)
	}
	productMap.RUnlock()
	sort.Ints(productIDs)
	return productIDs
}

func getNextProductID() int {
	productIDs := getProductIDs()
	return productIDs[len(productIDs)-1] + 1
}

func addOrUpdateProduct(product Product) (int, error) {
	// If the ProductID is set, Update, otherwise add
	addOrUpdateID := -1
	if product.ProductID > 0 {
		oldProduct := getProduct(product.ProductID)
		if oldProduct == nil {
			return 0, fmt.Errorf("Product id [%d] doesn't exist", product.ProductID)
		}
		addOrUpdateID = product.ProductID
	} else {
		addOrUpdateID = getNextProductID()
		product.ProductID = addOrUpdateID
	}
	productMap.Lock()
	productMap.m[addOrUpdateID] = product
	productMap.Unlock()
	return addOrUpdateID, nil
}
