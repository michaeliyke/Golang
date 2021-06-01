package main

import (
	"time"

	"example.com/log"
)

/*
Switch without a condition is the same as switch true.

This construct can be a clean way to write long if-then-else
 chains.
*/

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		log.P("Good morning!")
	case t.Hour() < 17:
		log.P("Good afternoon.")
	default:
		log.P("Good evening.")
	}
}
