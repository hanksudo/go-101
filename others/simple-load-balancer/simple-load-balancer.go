package main

import (
	"fmt"
	"time"
)

func sendLotsOfWork(in chan *Work) {

}

func receiveLotsOfResults(out chan *Work) {

}

// Work - Data struct
type Work struct {
	x, y, z int
}

func worker(in <-chan *Work, out chan<- *Work) {
	for w := range in {
		w.z = w.x * w.y
		fmt.Println(w.z)
		time.Sleep(time.Duration(w.z))
		out <- w
	}
}

func main() {
	in, out := make(chan *Work), make(chan *Work)
	for i := 0; i < 10; i++ {
		go worker(in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResults(out)
}
