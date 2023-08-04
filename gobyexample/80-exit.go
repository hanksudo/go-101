package main

import (
	"fmt"
	"os"
)

func main() {
	// won't be run
	defer fmt.Println("!")

	os.Exit(3)
}
