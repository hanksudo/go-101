package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	buffer := make([]byte, 8)
	for {
		n, err := r.Read(buffer)
		fmt.Printf("n=%v err=%v buffer=%v\n", n, err, buffer)
		fmt.Printf("buffer[:n] = %q\n", buffer[:n])
		if err == io.EOF {
			fmt.Println("Done!")
			break
		}
	}
}
