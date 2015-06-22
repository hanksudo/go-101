package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "I'm" // array is 0-based
	a[1] = " Hank"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var b [4]int // an array of type int
	fmt.Println("b ==", b)
	c := [10]int{1, 2, 3}
	d := [...]int{4, 5, 6}
	fmt.Println("c ==", c)
	fmt.Println("d ==", d)

	// two-dimensional arra
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	doubleSimple := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println("doubleArray==", doubleArray)
	fmt.Println("doubleSimple==", doubleSimple)

	// slices
	var ar = [5]byte{'a', 'b', 'c', 'd', 'e'} // define an array
	var ar1, ar2 []byte                       // define two slices with type []byte
	ar1 = ar[2:5]
	ar2 = ar[3:5]
	fmt.Println("ar1 ==", ar1)
	fmt.Println("ar2 ==", ar2)

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("slice ==", slice)
	fmt.Println("slice[:] ==", slice[:])
	fmt.Println("slice[0:7] ==", slice[0:7])
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%d] == %d \n", i, slice[i])
	}
	fmt.Println("slice[1:4] ==", slice[1:4]) // [3 4 5]
	fmt.Println("slice[:3] ==", slice[:3])   // [3 3 4]
	fmt.Println("slice[4:] ==", slice[4:])   // [6 7 8]
	fmt.Println("slice[4:7] ==", slice[4:7]) // [6 7 8]
	fmt.Println("slice[2:2] ==", slice[2:2]) // []

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

	// zero value of a slice is nil
	var t []int
	printSlice("t ==", t)
	fmt.Println(t == nil)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
