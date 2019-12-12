package main

import "math"

type Shape interface {
	Area() float64
}

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle -
type Circle struct {
	Radius float64
}

// Area returns the area of a circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle -
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of a triangle.
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
