// Declare a greeting package to collect related functions
package greetings

import (
	"fmt"
	"errors"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	// Declare and initialize a variable in one line
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string {
		"Hi, %v welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}