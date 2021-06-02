package main

import "example.com/log"

/*
Map literals are like struct literals, but the keys are required
*/

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{40.68433, -74.39967},
	"Google":    Vertex{37.42202, -122.08408},
}

func main() {
	log.P(m)
}
