package main

import (
	"fmt"
	"os"
	// "runtime"
)

func main() {
	// _, path, _, _ := runtime.Caller(0)
	path := "not-exist.go"
	if info, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s does not exists\n", path)
	} else {
		fmt.Println(info.Name())
		fmt.Println(info.Size())
		fmt.Println(info.Mode())
		fmt.Println(info.ModTime())
		fmt.Println(info.IsDir())
		// fmt.Println(info.Sys())
	}
}
