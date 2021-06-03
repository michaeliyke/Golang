package main

import (
  "math"

  "example.com/log"
)

/*
Functions are values too. They can be passed around just like
 other values.

Function values may be used as function arguments and return
 values
*/

func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

func main() {
  hypothenus := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }

  log.P(hypothenus(5, 12))
  log.P(compute(math.Pow))
}
