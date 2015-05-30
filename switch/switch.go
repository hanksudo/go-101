package main

import "fmt"

func main() {
	x := 5
	switch x {
	case 1:
		fmt.Println("x is equal to 1")
	case 3, 4, 5:
		fmt.Println("x is equal to 3, 4 or 5")
	default:
		fmt.Println("no matched")
	}

	y := 6
	switch y {
	case 4:
		fmt.Println("y is equal to 4")
		fallthrough
	case 5:
		fmt.Println("y is equal to 5")
		fallthrough
	case 6:
		fmt.Println("y is equal to 6")
		fallthrough
	case 7:
		fmt.Println("y is equal to 7")
		fallthrough
	case 8:
		fmt.Println("y is equal to 8")
		fallthrough
	default:
		fmt.Println("no matched")
	}
}
