package main

import (
	"fmt"
	"reflect"

	"golang.org/x/tour/tree"
)

//type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walk func(t *tree.Tree)
	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	trNums1 := make([]int, 0)
	trNums2 := make([]int, 0)

	ch := make(chan int)

	go Walk(t1, ch)
	for n := range ch {
		trNums1 = append(trNums1, n)
	}

	ch = make(chan int)

	go Walk(t2, ch)
	for n := range ch {
		trNums2 = append(trNums2, n)
	}

	return reflect.DeepEqual(trNums1, trNums2)
}

func main() {
	fmt.Println(Same(tree.New(4), tree.New(4)))
}
