package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex
var n = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var o = map[string]Vertex{
	"Bell Labes": {40.68433, -74.39967},
	"Google":     {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(n)
	fmt.Println(o)

	// --
	numbers := make(map[string]int)

	numbers["one"] = 1
	numbers["three"] = 3
	numbers["ten"] = 10
	fmt.Println("The first number is:", numbers["one"])
	fmt.Println("The third number is:", numbers["three"])

	delete(numbers, "ten")
	fmt.Println("The ten:", numbers["ten"])

	key, value := numbers["three"]
	fmt.Println("The key:", key, "value:", value)
}
