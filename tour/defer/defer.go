package main

import "fmt"

func add() (x int) {
	defer func() { x++ }()
	return 1
}
func main() {
	// defer
	defer fmt.Println("World")
	fmt.Println("Hello")

	// stacking defers (last-in-first-out)
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")

	fmt.Println(add()) // 2
}
