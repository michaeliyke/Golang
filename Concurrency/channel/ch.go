package main

import (
	"example.com/log"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1) // Buffered channel takes size arg to indicate how many msgs can wait
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		log.L(<-ch)
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 27 // This will wait to no avail for receiver
		wg.Done()
	}(ch, wg)
	wg.Wait()
}