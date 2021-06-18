package main

import (
	"math"

	"github.com/michaeliyke/Golang/log"
)

/*
Remember: a method is just a function with a receiver argument.

Here's Abs written as a regular function with no change in functionality.
*/

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	log.Println(Abs(v))
}
