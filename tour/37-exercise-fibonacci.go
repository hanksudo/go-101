// Exercise: Fibonacci closure
// https://tour.golang.org/moretypes/26
package main

import (
	"fmt"
)

func fibonacci() func() int {
	first, next := 0, 1
	return func() int {
		first, next = next, first+next
		return next - first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
