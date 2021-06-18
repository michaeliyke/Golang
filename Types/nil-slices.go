package main

import "github.com/michaeliyke/Golang/log"

/*
The zero value of a slice is nil.
A nil slice has a length and capacity of 0 and has no underlying
 array.
*/

func main() {
	var s []int
	P(s)

	if s == nil {
		log.P("Nil")
	}
}

func P(x []int) {
	log.Print(x)
	log.Print(":- ")
	log.Print("len - ")
	log.Print(len(x))
	log.Print(" cap - ")
	log.Println(cap(x))
}
