package main

import "fmt"

func main() {
	// The type []T is a slice with elements of type T.
	primes := [6]int{2, 3, 5, 7, 11, 13}

	s := primes[1:4]
	fmt.Println(s)
	fmt.Println()

	// define an array
	var ar = [5]byte{'a', 'b', 'c', 'd', 'e'}

	// define two slices with type []byte
	var ar1, ar2 []byte
	ar1 = ar[2:5]
	ar2 = ar[3:5]
	fmt.Println("ar1 ==", ar1)
	fmt.Println("ar2 ==", ar2)
	fmt.Println()

	// 'make' function - slices created
	x := make([]int, 5)
	printSlice("x", x)
	y := make([]int, 0, 5)
	printSlice("y", y)
	z := y[:2]
	printSlice("z ==", z)
	r := z[2:5]
	printSlice("r ==", r)
	// exceeds capacity, x is reallocated
	rr := append(x, 1, 2)
	printSlice("rr ==", rr)
	fmt.Println()

	// append slice to slice
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s1 = append(s1, s2...)
	fmt.Println("s1 ==", s1)
	fmt.Println()

	// zero value of a slice is nil
	var t []int
	printSlice("t ==", t)
	fmt.Println(t == nil)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
