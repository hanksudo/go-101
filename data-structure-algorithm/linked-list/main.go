package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}
type LinkedList struct {
	head *Node
}

func (l *LinkedList) Prepend(n *Node) {
	n.next = l.head
	l.head = n
}

func (l *LinkedList) Insert(n *Node) {
	if l.head == nil {
		l.head = n
		return
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = n
}

func (l *LinkedList) DeleteByValue(value int) {
	if l.head == nil {
		return
	}
	if l.head.value == value {
		l.head = l.head.next
		return
	}

	node := l.head
	for node.next != nil {
		if node.next.value == value {
			node.next = node.next.next
			return
		}
		node = node.next
	}
}

func (l LinkedList) PrintList() {
	node := l.head
	for node != nil {
		fmt.Printf("%d ", node.value)
		node = node.next
	}
	fmt.Println()
}

func main() {
	l := LinkedList{}
	l.Prepend(&Node{value: 1})
	l.Prepend(&Node{value: 3})
	l.Prepend(&Node{value: 5})
	l.PrintList()

	l.Insert(&Node{value: 7})
	l.Insert(&Node{value: 9})
	l.PrintList()

	l.DeleteByValue(7)
	l.PrintList()

	emptyList := LinkedList{}
	emptyList.DeleteByValue(10)
}
