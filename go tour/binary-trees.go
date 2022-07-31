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
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func treeLen(t *tree.Tree) int {
	if t == nil {
		return 0
	}
	if t.Left == nil && t.Right == nil {
		return 1
	}
	return 1 + treeLen(t.Left) + treeLen(t.Right)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	trNums1 := make([]int, 0)
	trNums2 := make([]int, 0)

	tl := treeLen(t1)
	ch := make(chan int, tl)

	go Walk(t1, ch)
	for i := 0; i < tl; i++{
		n := <-ch
		trNums1 = append(trNums1, n)
	}

	tl = treeLen(t2)
	ch = make(chan int, tl)

	go Walk(t2, ch)
	for i := 0; i < tl; i++{
		n := <-ch
		trNums2 = append(trNums2, n)
	}

	return reflect.DeepEqual(trNums1, trNums2)
}

func main() {
	fmt.Println(Same(tree.New(4), tree.New(2)))
}
