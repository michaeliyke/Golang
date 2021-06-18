package main

import "github.com/michaeliyke/Golang/log"

/*
The type [n]T is an array of n values of type T.

The expression

var a [10]int
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized.
This seems limiting, but don't worry; Go provides a
convenient way of working with array
*/

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = " world"
	log.P(a[0] + a[1])

	primes := [6]int{2, 3, 5, 7, 11, 13}
	log.P(primes)
}
