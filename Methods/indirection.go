package main

import (
	"math"

	"github.com/michaeliyke/Golang/log"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	log.Println(Abs(v))

	v.Scale(2)
	log.Println(Abs(v))

	ScaleFunc(&v, 5)
	log.Println(Abs(v))
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.Y = v.Y * f
	v.X = v.X * f
}
