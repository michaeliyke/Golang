package main

import "github.com/michaeliyke/Golang/log"

/*
Struct fields are accessed using a dot.
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{20, 30}
	v.X = 52
	log.P(v.X)
}
