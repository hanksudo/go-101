// A Tour of Go Exercise
// https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
	"math"
)

// Sqrt - function to return sqrt
func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for math.Abs(z-tmp) > 1e-7 {
		tmp = z
		z -= (math.Pow(z, 2) - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(3))
	fmt.Println(math.Sqrt(3))
	fmt.Println(Sqrt(5))
	fmt.Println(math.Sqrt(5))
}
