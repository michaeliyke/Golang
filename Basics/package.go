package main

import (
	"fmt"
)

func main() {
	fmt.Println(split(17))
}

// Go's return values may be named, eg x int, y int.
// If so, they are treated as variables defined at the top of the function
// A return statement without arguments returns the named return values.
// This is known as a "naked" return.
func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
