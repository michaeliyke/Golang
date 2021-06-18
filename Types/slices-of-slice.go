package main

import (
	"strings"

	"github.com/michaeliyke/Golang/log"
)

/*
Slices can contain any type, including other slices.
*/

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		s := strings.Join(board[i], " ")
		log.P(s)
	}
}

func P(x []int, s string) {
	log.Print(s + ":- ")
	log.Print("len - ")
	log.Print(len(x))
	log.Print(" cap - ")
	log.Print(cap(x))
	log.Print(" value - ")
	log.Println(x)
}
