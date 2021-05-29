/*
Glass() - Glass returns a useful phrase for world travelers.
Go() - Go returns a Go proverb.
Hello() - Hello returns a greeting.
Opt() - Opt returns an optimization truth.
*/

package main

import "fmt"
import "rsc.io/quote"

func main() {
	const x = 40
	var y = 10
	// var random = rand(12000)
	fmt.Println("Hello world!", x, y)
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
