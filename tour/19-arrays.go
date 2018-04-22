package main

import (
	"fmt"
)

func main() {
	// Arrays - The type [n]T is an array of n values of type T.
	// arrays can't be resized.
	var a [2]string
	a[0] = "I'm" // array is 0-based
	a[1] = " Hank"
	// a[2] = "Hello" // out of bounds for 2-element array
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println()

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println()

	var b [4]int // an array of type int
	fmt.Println("b ==", b)

	// fill with zero
	c := [10]int{1, 2, 3}
	fmt.Println("c ==", c)

	d := [...]int{4, 5, 6, 7}
	fmt.Println("d ==", d)

	// two-dimensional array
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	doubleSimple := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println("doubleArray  ==", doubleArray)
	fmt.Println("doubleSimple ==", doubleSimple)
}
