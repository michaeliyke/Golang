package main

import "github.com/michaeliyke/Golang/log"

func main() {
	var s []int
	P(s, "[]int")

	// Append works on nil slices
	s = append(s, 0)
	P(s, "append(s, 0)")

	s = append(s, 1)
	P(s, "append(s, 1)")

	s = append(s, 2, 3, 4)
	P(s, "append(s, 2, 3, 4)")
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
