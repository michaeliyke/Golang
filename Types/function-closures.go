package main

import "github.com/michaeliyke/Golang/log"

/*
Function closures
Go functions may be closures. A closure is a function value that
 references variables from outside its body. The function may
 access and assign to the referenced variables; in this sense the
 function is "bound" to the variables.
For example, the adder function returns a closure. Each closure
is bound to its own sum variable.
*/

func adder() func(int) int {
	var sum int = 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	var positive, negative func(int) int = adder(), adder()
	for index := 0; index < 10; index++ {
		log.Println(positive(index), negative(-2*index))
	}
}
