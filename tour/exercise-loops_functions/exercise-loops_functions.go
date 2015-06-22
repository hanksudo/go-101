// A Tour of Go Exercise
// http://tour.golang.org/#25
package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	tmp := 0.0
	for math.Abs(z-tmp) > 1e-6 {
		tmp = z
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
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
