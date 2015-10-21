package main

import "fmt"

// only accepts a channel for sending
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// accepts one channel for receives and one for sends
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
