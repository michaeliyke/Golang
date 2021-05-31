// Programs start running in package main.
package main

// In Go, a name is exported if it begins with a capital letter.
// For example, Pizza is an exported name, as is Pi, which is exported from the math package

// When importing a package, you can refer only to its exported names.
// Any "unexported" names are not accessible from outside the package.
import (
	"fmt"
)

func main() {
	x, y := swap("Iyke", "Michael")
	fmt.Println(x, y)
}
func swap(x string, y string) (string, string) {
	return y, x
}
