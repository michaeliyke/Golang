package main

import "example.com/log"

func Sqrt(x float64) float64 {
	var j float64
	for z := float64(1); z*z < x+1; z++ {
		if z*z < x {
			j = z
		}
		if z*z == x {
			return z
		}
		if z*z > x {
			return j
		}
	}
	return j
}

func main() {
	log.Log(Sqrt(99))
}
