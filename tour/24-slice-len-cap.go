package main

import "fmt"

func main() {
	// The length of a slice is the number of elements.
	// The capacity of a slice is the number of elements in the underlying array.

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("", s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice("", s)

	// Extends its length.
	s = s[:4]
	printSlice("", s)

	// Drop its first two values
	s = s[2:]
	printSlice("", s)
	fmt.Println("")

	// zero value of a slice is nil
	var t []int
	printSlice("t ==", t)
	fmt.Println(t == nil)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
