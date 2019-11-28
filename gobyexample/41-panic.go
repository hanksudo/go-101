package main

import (
	"os"
)

func main() {
	// panic("a problem")
	_, err := os.Create("/tmp2/file")
	if err != nil {
		panic(err)
	}
}
