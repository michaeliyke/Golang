package main

import "fmt"

func fibbonacci() func(int) []int {
	current, placeholder, foo := 0, 0, 1
	var array []int

	var fib func(int) []int
	fib = func(number int) []int {
		current += foo
		array = append(array, placeholder)
		foo = placeholder
		placeholder = current
		if placeholder < number {
			return fib(number)
		}
		return array
	}

	return fib
}

func main() {
	// var fib []int = fibbonacci()
	f := fibbonacci()
	for index := 0; index < 10; index++ {
		fmt.Println(f(index))
	}
}
