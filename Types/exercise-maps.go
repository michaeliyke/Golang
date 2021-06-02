/*
Exercise: Maps
Implement WordCount. It should return a map of the counts of each
 “word” in the string s. The wc.Test function runs a test suite
 against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var words []string = strings.Fields(s)
	wordcount := make(map[string]int)
	for _, word := range words {
		// word = strings.ToLower(word)
		_, present := wordcount[word]
		if present == true {
			wordcount[word] += 1
		} else {
			wordcount[word] = 1
		}
	}
	return wordcount
}

func main() {
	// log.Println(WordCount("mad man is making some mad joke and is not smiling over it yet he is laughing"))
	wc.Test(WordCount)
}
