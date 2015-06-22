package main

import (
	"fmt"
)

var arr = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range arr {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	// discard certain return values
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
