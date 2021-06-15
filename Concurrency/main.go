package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	x := 0
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		id := rand.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				x++
				log.Println("################################", x)
				log.Println("From cache")
				log.Println(b)
			}
			wg.Done()
		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				x++
				log.Println("################################", x)
				log.Println("From Database")
				log.Println(b)
			}
			wg.Done()
		}(id, wg, m)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.RLock()
			cache[id] = b
			m.RUnlock()
			return b, true
		}
	}
	return Book{}, false
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.Lock()
	b, ok := cache[id]
	m.Unlock()
	return b, ok
}
