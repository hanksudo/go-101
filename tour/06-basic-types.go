package main

import (
	"fmt"
	"math/cmplx"
)

var (
	enabled    bool       = true
	disabled              = false
	MaxInt     uint64     = 1<<64 - 1
	z          complex128 = cmplx.Sqrt(-5 + 12i)
	c          complex64  = 5 + 5i
	runeNumber rune       = 5         // alias of `int32`
	byteNumber byte       = 5         // alias of `uint8`
	PI                    = 3.1415926 // default `float64`
)

func main() {
	const f = "%T(%v)\n"

	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	fmt.Printf(f, c, c)
	fmt.Printf(f, enabled, enabled)
	fmt.Printf(f, disabled, disabled)
	fmt.Printf(f, runeNumber, runeNumber)
	fmt.Printf(f, byteNumber, byteNumber)
	fmt.Printf(f, PI, PI)
}
