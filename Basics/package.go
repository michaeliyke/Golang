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
	fmt.Println("Sum: ", add(42, 13))
	fmt.Println("DIfference: ", subtract(42, 13))
}

// Notice that the type comes after the variable name
func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
func subtract(z, y int) int {
	return z - y
}
