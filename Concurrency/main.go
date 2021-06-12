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
		go func(id int) {
			if b, ok := queryCache(id); ok {
				log.Println("From cache")
				log.Println(b)
			}
		}(id)
		go func(id int) {
			if b, ok := queryDatabase(id); ok {
				log.Println("From Database")
				log.Println(b)
			}
		}(id)
		time.Sleep(150 * time.Millisecond)
	}
	time.Sleep(2 * time.Second)
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
