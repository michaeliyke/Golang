package main

import "github.com/michaeliyke/Golang/log"

/*
Slices can be created with the built-in make function; this is
how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice
 that refers to that array:

a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/

func main() {
	a := make([]int, 5)
	P(a, "make([]int, 5)")

	b := make([]int, 0, 5)
	P(b, "make([]int, 0, 5)")

	c := b[:2]
	P(c, "b[:2]")

	d := c[2:5]
	P(d, "c[2:5]")
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
