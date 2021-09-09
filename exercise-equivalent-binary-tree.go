package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){	
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

// hàm so sánh
func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}

		if !ok1{
			break
		}
	}
	return true
}

func main() {
	// go Walk(tree.New(1), ch)
	// tree.New : trả về 1 cây nhị phân ngẫu nhiên
	fmt.Println(Same(tree.New(2), tree.New(1)))
}
