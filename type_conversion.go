package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y int = 3, 4 // x, y := 3, 4
    var f float64 = math.Sqrt(float64(3*3 + 4*4)) // f := math.Sqrt(float64(3*3 + 4*4))
    var z int = int(f) // z := int(f)
    fmt.Println(x, f, y, z)
}