package main

import "example.com/log"

/*
A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the
underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using
the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has
sufficient capacity. Try changing one of the slice operations in
the example program to extend it beyond its capacity and see what
happens.
*/

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[:0]
	P(s, "s[:0]")

	s = s[:4]
	P(s, "s[:4]")

	s = s[2:]
	P(s, "s[2:]")
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
