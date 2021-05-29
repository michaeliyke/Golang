// Declare a greeting package to collect related functions
package greetings

import "fmt"

func Hello(name string) string {
	// Declare and initialize a variable in one line
	message := fmt.Sprintf("Hi, %v. welcome", name)
	return message
}
