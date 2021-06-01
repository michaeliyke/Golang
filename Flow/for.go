package main

import "example.com/log"

/*
A defer statement defers the execution of a function until t
he surrounding function returns.

The deferred call's arguments are evaluated immediately, but
 the function call is not executed until the surrounding function
 returns.
*/

func main() {
	defer log.P("world")

	log.P("hello")
}
