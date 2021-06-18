package main

import "github.com/michaeliyke/Golang/log"

/*

Slice defaults

When slicing, you may omit the high or low bounds to use their
defaults instead.

The default is zero for the low bound and the length of the slice
for the high bound.

For the array

var a [10]int
these slice expressions are equivalent:

a[0:10]
a[:10]
a[0:]
a[:]

*/

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	s1 := s[1:4]
	log.P(s1)

	s2 := s[:2]
	log.P(s2)
	s3 := s[1:]
	log.P(s3)

}
