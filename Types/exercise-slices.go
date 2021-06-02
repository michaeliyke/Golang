package main

import (
	"golang.org/x/tour/pic"
)

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy) //slice of slices of length dy
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			// s[x] = uint8((x + y) / 2)
			// s[x] = uint8((x * y) / 2)
			s[x] = uint8((x ^ y))
			// s[x] = uint8((x ^ y) / 2)
			// s[x] = uint8((x + y))
		}
		ss[y] = s
	}
	return ss
}
