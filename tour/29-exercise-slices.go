// Exercise: Slices
// https://tour.golang.org/moretypes/18

package main

import (
	"code.google.com/p/go-tour/pic"
)

// Pic ...
func Pic(dx, dy int) [][]uint8 {
	// return slice of length dy
	ret := make([][]uint8, dy)

	// each element of which is a slice
	// of dx 8-bit unsigned integers.
	for y := range ret {
		ret[y] = make([]uint8, dx)

		// The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
		for x := range ret[y] {
			ret[y][x] = uint8(x ^ y)
		}
	}

	return ret
}

func main() {
	pic.Show(Pic)
}
