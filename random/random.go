// http://golang.org/pkg/math/rand/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My lucky number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))
	fmt.Println(math.Pi)
}
