package main

import (
	"fmt"

	"github.com/michaeliyke/Golang/log"
)

/*
The range form of the for loop iterates over a slice or map.

When ranging over a slice, two values are returned for each
iteration. The first is the index, and the second is a copy of
the element at that index.
*/

func main() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func P(x []int, s string) {
	log.Print(s + ":- ")
	log.Print("len - ")
	log.Print(len(x))
	log.Print(" cap - ")
	log.Print(cap(x))
	log.Print(" value - ")
	log.Println(x)
}
