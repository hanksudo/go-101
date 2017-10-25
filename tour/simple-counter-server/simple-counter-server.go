package main

import (
	"fmt"
	"log"
	"net/http"
)

type counter struct {
	n int
}

func (ctr *counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func main() {
	ctr := new(counter)
	http.Handle("/counter", ctr)
	err := http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
