package main

import (
	"math"

	"example.com/log"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	log.Println(Abs(v))
	Scale(&v, 10)
	log.Println(Abs(v))
}
