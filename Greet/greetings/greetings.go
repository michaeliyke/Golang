// Declare a greeting package to collect related functions
package greetings

import (
	"fmt"
	"errors"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	// Declare and initialize a variable in one line
	message := fmt.Sprintf("Hi, %v. welcome", name)
	return message, nil
}
