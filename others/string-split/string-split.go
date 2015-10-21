package main

import (
	"fmt"
	"strings"
)

func main() {
	// tempArr := strings.Split("1 3 5 7 9", " ")
	var tempArr = strings.Split("1 3 5 7 9", " ")
	for i := 0; i < len(tempArr); i++ {
		fmt.Println(tempArr[i])
	}

}
