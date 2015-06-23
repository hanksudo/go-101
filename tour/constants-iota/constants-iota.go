// iota make enum, begin with 0, increased by 1
package main

import (
	"fmt"
)

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w        // If there is no expression, it use the last expression, so w = iota = 3
)

// once iota meets keyword `const`, it reset to 0, so v = 0
const v = iota

const (
	a, b, c = iota, iota, iota // values of iota are same in one line.
	d, e, f = iota, iota, iota
)

func main() {
	fmt.Println(x, y, z, w)
	fmt.Println(v)
	fmt.Println(a, b, c, d, e, f)
}
