package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(n int) {
	q.items = append(q.items, n)
}

func (q *Queue) Dequeue() int {
	value := q.items[0]
	q.items = q.items[1:]
	return value
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(1)
	myQueue.Enqueue(3)
	myQueue.Enqueue(5)
	fmt.Println(myQueue)

	myQueue.Dequeue()
	fmt.Println(myQueue)
}
