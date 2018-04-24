package main

import (
	"fmt"
)

func main() {
	// Empty interfaces are used by code that
	// handles values of unknown type.
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
