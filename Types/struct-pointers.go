package main

import "example.com/log"

/*
To access the field X of a struct when we have the struct pointer
 p we could write (*p).X. However, that notation is
 cumbersome, so the language permits us instead to write
 just p.X, without the explicit dereference.
*/

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// p := &v
	p.X = 1e9
	(*p).Y = 1e6
	log.P(v)
}
