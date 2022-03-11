package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) Pop() int {
	index := len(s.items) - 1
	lastElement := s.items[index]
	s.items = s.items[:index]
	return lastElement
}

func main() {
	myStack := Stack{}
	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)
	myStack.Push(5)
	fmt.Println(myStack)

	myStack.Pop()
	fmt.Println(myStack)
}
