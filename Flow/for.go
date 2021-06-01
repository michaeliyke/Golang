package main

import "example.com/log"

/*
Deferred function calls are pushed onto a stack. When a
function returns, its deferred calls are executed in
last-in-first-out order.
https://blog.golang.org/defer-panic-and-recover
*/

func main() {
	log.P("counting")

	for i := 0; i <= 10; i++ {
		defer log.P(i)
	}

	log.P("done")
}
