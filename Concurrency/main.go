package main

import (
	"log"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rand.Intn(10) + 1
		if b, ok := queryCache(id); ok {
			log.Println("From cache")
			log.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			log.Println("From Database")
			log.Println(b)
			continue
		}
		log.Printf("Book not found with id: %v", id)
		time.Sleep(150 * time.Millisecond)
	}
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}
