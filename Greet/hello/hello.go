// Declare a greeting package to collect related functions
package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	// Declare and initialize a variable in one line
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
