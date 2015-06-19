package main

import (
	"fmt"
)

func fibonacci() func() int {
	f, g := 0, 1
	return func() int {
		f, g = g, f+g
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
