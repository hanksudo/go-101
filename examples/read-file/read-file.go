package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var result []byte
	buf := make([]byte, 1)
	for {
		n, err := f.Read(buf)
		result = append(result, buf[:n]...)
		fmt.Printf("n=%v err=%v buf=%v\n", n, err, buf)
		fmt.Printf("buf[:n] = %q\n", buf[:n])
		if err == io.EOF {
			fmt.Println("Done!")
			break
		}
	}
	fmt.Println("result=", string(result))
}
