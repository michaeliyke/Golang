// Declare a greeting package to collect related functions
package main

import (
	"fmt"
	"example.com/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Declare and initialize a variable in one line
	message, error := greetings.Hello("")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(message)
}
