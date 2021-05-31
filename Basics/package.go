// Programs start running in package main.
package main

// By convention, the package name is the same as the last element of the import path
// the "math/rand" package comprises files that begin with the statement package rand
import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Printf("Now you have %g problems \n", math.Sqrt(7))
	fmt.Println("My favorite number is ", rand.Intn(10))
}
