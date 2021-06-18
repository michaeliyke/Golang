package main

import (
	"github.com/michaeliyke/Golang/log"
)

/*
Mutating Maps
Insert or update an element in map m:

m[key] = elem
Retrieve an element:

elem = m[key]
Delete an element:

delete(m, key)
Test that a key is present with a two-value assignment:

elem, ok = m[key]
If key is in m, ok is true. If not, ok is false.

If key is not in the map, then elem is the zero value for the map's element type.

Note: If elem or ok have not yet been declared you could use a short declaration form:

elem, ok := m[key]
*/
func main() {
	m := make(map[string]int)
	m["Answer"] = 42
	log.P("The value ", m["Answer"])

	m["Answer"] = 48
	log.P("The value ", m["Answer"])

	delete(m, "Answer")
	log.P("The value ", m["Answer"])

	v, ok := m["Answer"]
	log.P("The value ", v, "Present?", ok)
}
