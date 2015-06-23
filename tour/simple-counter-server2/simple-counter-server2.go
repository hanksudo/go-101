package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Counter int
type Chan chan *http.Request

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}

// Argument Server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func main() {
	ctr := new(Counter)
	http.Handle("/counter", ctr)
	http.Handle("/args", http.HandlerFunc(ArgServer))

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
