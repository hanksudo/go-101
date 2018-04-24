package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walker(t *tree.Tree, ch, quit chan int) {
	if t == nil {
		return
	}
	// fmt.Println(t)
	// fmt.Println("Walker:", t.Value)
	Walker(t.Left, ch, quit)
	select {
	case ch <- t.Value:
		// Value successfully sent.
	case <-quit:
		return
	}
	Walker(t.Right, ch, quit)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch
func Walk(t *tree.Tree, ch, quit chan int) {
	Walker(t, ch, quit)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	quit := make(chan int)
	defer close(quit)

	go Walk(t1, ch1, quit)
	go Walk(t2, ch2, quit)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	// ch := make(chan int)
	// go Walk(tree.New(1), ch)

	// for i := range ch {
	// 	fmt.Println(i)
	// }

	fmt.Println("tree.New(1) == tree.New(1):", Same(tree.New(1), tree.New(1)))
	fmt.Println("tree.New(1) == tree.New(2):", Same(tree.New(1), tree.New(2)))
}
