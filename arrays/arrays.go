package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "I'm"
	a[1] = " Hank"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var b [4]int
	fmt.Println(b)

	// slices
	p := []int{3, 3, 4, 5, 6, 7, 8}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d \n", i, p[i])
	}

	fmt.Println("p[1:4] ==", p[1:4]) // [3 4 5]
	fmt.Println(p[:])
	fmt.Println(p)
	fmt.Println("p[:3] ==", p[:3])   // [3 3 4]
	fmt.Println("p[4:] ==", p[4:])   // [6 7 8]
	fmt.Println("p[2:2] ==", p[2:2]) // []

	// 'make' function - slices created
	x := make([]int, 5)
	printSlice("x", x)
	y := make([]int, 0, 5)
	printSlice("y", y)
	z := y[:2]
	printSlice("z", z)
	r := z[2:5]
	printSlice("r", r)

	// zero value of a slice is nil
	var t []int
	printSlice("t", t)
	fmt.Println(t == nil)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
