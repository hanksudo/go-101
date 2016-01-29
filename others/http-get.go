package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.instagram.com/p/BACozJ6kUhc/?taken-by=hanksudo")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(contents))
}
