package main

import "fmt"

func main() {
	var s []int
	printSlice(s)
	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slices grows as needed
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)

	// append slice to slice
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	printSlice(s1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}
