package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type counter int
type _chan chan *http.Request

func (ctr *counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

func (ch _chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}

// Argument Server.
func argServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func headerHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, req.Header)
}

func main() {
	ctr := new(counter)
	http.Handle("/counter", ctr)
	http.Handle("/args", http.HandlerFunc(argServer))
	http.Handle("/header", http.HandlerFunc(headerHandler))

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
