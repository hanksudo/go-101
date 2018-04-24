package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walker(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	// fmt.Println(t)
	// fmt.Println("Walker:", t.Value)
	Walker(t.Left, ch)
	ch <- t.Value
	Walker(t.Right, ch)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch
func Walk(t *tree.Tree, ch chan int) {
	Walker(t, ch)
	// Need to close channel here
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// !!!: The implementation leaks goroutines when trees are different.
// see binarytrees_quit.go for a better solution.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

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
