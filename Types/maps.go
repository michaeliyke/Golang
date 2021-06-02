package main

import "example.com/log"

/*
A map maps keys to values.

The zero value of a map is nil. A nil map has no keys, nor c
an keys be added.

The make function returns a map of the given type, initialized
and ready for use.
*/

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.68433, -7439967}
	log.Println(m["Bell Labs"])
}
