package main

import (
	"fmt"
)

func main() {
	// sum
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum:", sum)

	// sum2 - identical to while
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println("sum2:", sum2)

	// break & continue
	for i := 0; i < 10; i++ {
		if i == 3 { // skip
			continue
		}
		if i > 5 { // terminate
			break
		}
		fmt.Println(i)
	}

	// Multiplication
	for i := 2; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Printf("%d x %d = %d \n", i, j, i*j)
		}
		fmt.Println()
	}

	// forever
	// for {
	// }
}
