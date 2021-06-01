package main

import (
	"time"

	"example.com/log"
)

func main() {
	log.P("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		log.P("Today.")
	case today + 1:
		log.P("Tomorrow.")
	case today + 2:
		log.P("In two days.")
	default:
		log.P("Too far away.")
	}
}
