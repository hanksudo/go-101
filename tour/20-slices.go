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
}
