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

func HelloError(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	// Declare and initialize a variable in one line
	message := fmt.Sprintf(randomFormat())
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
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