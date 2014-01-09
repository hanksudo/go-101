package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Happy", math.Pi, "Day")
    fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}

