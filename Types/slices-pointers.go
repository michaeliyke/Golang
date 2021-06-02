package main

import "example.com/log"

/*
A slice does not store any data, it just describes a section of
an underlying array.

Changing the elements of a slice modifies the corresponding
elements of its underlying array.

Other slices that share the same underlying array will see those
 changes.
*/

func main() {
	names := [4]string{
		"John", "Paul", "George", "Ringo",
	}
	log.P(names)
	a := names[0:2]
	b := names[2:4]
	log.Println(a)
	log.Println(b)
	b = names[1:3]
	b[0] = "XXX"

	log.P("\n:::::::::::::::")
	log.P("AFTER MODIFICATION")
	log.P(":::::::::::::::\n")
	log.Println(a)
	log.Println(b)
	log.P(names)
}
