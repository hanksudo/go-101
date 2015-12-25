package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"strings"

	"fmt"
)

func main() {
	dat, err := ioutil.ReadFile("./test.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(strings.NewReader(string(dat)))
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatalf("error reading all lines: %v", err)
	}

	for i, line := range lines {
		if i == 0 {
			continue
		}
		fmt.Println(line[0], line[1])
	}
}
