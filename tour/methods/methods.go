// http://tour.golang.org/methods/1
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
