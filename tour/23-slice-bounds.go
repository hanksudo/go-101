package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	// These slice expressions are equivalent
	// fmt.Println(s[0:6])
	// fmt.Println(s[:6])
	// fmt.Println(s[0:])
	// fmt.Println(s[:])

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("slice      ==", slice)
	fmt.Println("slice[:]   ==", slice[:])
	fmt.Println("slice[0:7] ==", slice[0:7])
	fmt.Println()

	// iterate slices
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] == %d \n", i, slice[i])
	}
	fmt.Println()

	fmt.Println("slice[1:4] ==", slice[1:4]) // [3 4 5]
	fmt.Println("slice[:3]  ==", slice[:3])  // [3 3 4]
	fmt.Println("slice[4:]  ==", slice[4:])  // [6 7 8]
	fmt.Println("slice[4:7] ==", slice[4:7]) // [6 7 8]
	fmt.Println("slice[2:2] ==", slice[2:2]) // []
	fmt.Println()
}
