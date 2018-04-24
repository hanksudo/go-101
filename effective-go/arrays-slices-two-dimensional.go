// https://golang.org/doc/effective_go.html#two_dimensional_slices
package main

import (
	"fmt"
)

type Transform [3][3]float64 // A 3x3 array, really an array of arrays
type LinesOfText [][]byte    // A slice of byte slices.

func main() {
	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}

	fmt.Println(text)

	t := Transform{}
	fmt.Println(t)
}
