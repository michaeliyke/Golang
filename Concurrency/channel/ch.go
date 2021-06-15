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
		//if value, ok := ch; ok{}
		for v := range ch {
			log.Log(v)
		}
		wg.Done()
	}(ch, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}(ch, wg)
	wg.Wait()
}