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
	names := []string{"Gladys", "Samantha", "Darrin"}
	// Declare and initialize a variable in one line
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
