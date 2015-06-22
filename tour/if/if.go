package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x)
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}

func gotoFunc() {
	i := 0
Here:
	fmt.Println(i)
	i++
	if i < 10 {
		goto Here
	}
}

func main() {
	x := 6
	if x%2 == 0 {
		fmt.Println("x is even number.")
	} else {
		fmt.Println("x is odd number.")
	}

	// initialize y in if block
	if y := 3; y%2 == 0 {
		fmt.Println("y is even number.")
	} else {
		fmt.Println("y is odd number.")
	}

	fmt.Println(
		"sqrt(2) ==",
		sqrt(2),
		", sqrt(-4) ==",
		sqrt(-4),
	)

	gotoFunc()

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
