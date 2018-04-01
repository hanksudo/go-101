package main

import "fmt"

// (x int, y int) = (x, y int)
func add(x int, y int) int {
	return x + y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Multi results
func swap(x, y string) (string, string) {
	return y, x
}

// Named results - will return x, y
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Variable arguments
func list(args ...int) {
	for _, n := range args {
		fmt.Printf("number is %d\n", n)
	}
}

// pass by pointer
func add1(x *int) {
	*x = *x + 1
	return
}

func main() {
	fmt.Printf("add(%d, %d) = %d\n", 42, 13, add(42, 13))
	fmt.Printf("max(%d, %d) = %d\n", 42, 13, max(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	list(1, 2, 3)

	x := 5
	add1(&x)
	fmt.Println("x =", x)

	// defer - execute by reverse order
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
		fmt.Printf("%d ", i)
	}
}
