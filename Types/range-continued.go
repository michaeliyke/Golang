package main

import (
	"fmt"

	"github.com/michaeliyke/Golang/log"
)

/*
You can skip the index or value by assigning to _.

for i, _ := range pow
for _, value := range pow
If you only want the index, you can omit the second variable.

for i := range pow
*/

func main() {
	var pow = make([]int, 10)
	for index := range pow {
		pow[index] = 1 << uint(index) //2**i
	}
	for index, value := range pow {
		fmt.Printf("%d: ", index+1)
		fmt.Printf("%d\n", value)
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
