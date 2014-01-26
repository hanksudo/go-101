package main

import (
	"fmt"
)

func main() {

	// Multiplication
	for i := 2; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d x %d = %d \n", i, j, i*j)
		}
		fmt.Println()
	}

	// sum
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// sum2
	sum2 := 1
	for sum2 <= 100 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// forever
	// for {
	// }
}
