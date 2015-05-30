package main

import "fmt"

type testInt func(int) bool // define a function type of variable

func isOdd(n int) bool {
	return n%2 != 0
}

func isEven(n int) bool {
	return n%2 == 0
}

// pass the function `f` as an argument to another function
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slice =", slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd numbers =", odd)
	even := filter(slice, isEven)
	fmt.Println("even numbers =", even)
}
