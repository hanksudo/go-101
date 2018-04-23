package main

import "fmt"

func main() {
	// 'make' function - slices created
	a := make([]int, 5)
	b := make([]int, 0, 5)
	c := b[:2]
	d := c[2:5]
	printSlice("a ==", a)
	printSlice("b ==", b)
	printSlice("c ==", c)
	printSlice("d ==", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
