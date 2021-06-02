package main

import "example.com/log"

/*
   Go has pointers. A pointer holds the memory address of a value.
   The * operator denotes the pointer's underlying value.
   Unlike C, Go has no pointer arithmetic.
*/

func main() {
	var i int = 42

	// The type *T is a pointer to a T value. Its zero value is nil.
	var p *int

	// The & operator generates a pointer to its operand.
	// This is known as "dereferencing" or "indirecting".
	p = &i
	log.Log(p)
	log.P(i)

	// This is known as "dereferencing" or "indirecting".
	*p = 21
	log.P(p)
	log.P(i)
}
