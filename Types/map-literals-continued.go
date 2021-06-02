package main

import "example.com/log"

/*

If the top-level type is just a type name, you can omit it from
the elements of the literal.
*/

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	log.P(m)
}
