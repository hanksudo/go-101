package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Happy", math.Pi, "Day")
	fmt.Printf("Now you have %g problems. \n", math.Nextafter(2, 3))
	fmt.Println("1 << 10 :", 1<<10) // 1 * 1 ^ 10
	fmt.Println("2 << 10 :", 2<<10) // 2 * 2 ^ 10
	fmt.Println("8 >> 2:", 8>>2)    // 8 * 2 ^ -2 or 8 / 2 ^ 2
	fmt.Println("math.Pow(2, 10)", math.Pow(2, 10))
}
