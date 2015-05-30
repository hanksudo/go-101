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

	// ---
	p := make(map[string]int)

	p["Answer"] = 42
	fmt.Println("The value:", p["Answer"])

	p["Answer"] = 48
	fmt.Println("The value:", p["Answer"])

	delete(p, "Answer")
	fmt.Println("The value:", p["Answer"])

	v, ok := p["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
