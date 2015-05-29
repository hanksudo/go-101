package main

import (
	"fmt"
)

const (
	Big   = 1 << 100  // 1 * 2 ^ 100
	Small = Big >> 99 // Big * 2 ^ -99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big)) // overflows
	fmt.Println(needFloat(Big))
	fmt.Println(needFloat(Big))
}