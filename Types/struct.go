package main

import "example.com/log"

/*
A struct is a collection of fields.
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	log.P(Vertex{1, 2})
}
