// package main

// import (
// 	"fmt"

// 	"golang.org/x/tour/tree"
// )

// // Walk walks the tree t sending all values
// // from the tree to the channel ch.
// func Walk(t *tree.Tree, ch chan int) {
// 	var walker func(node *tree.Tree, ch chan int)
// 	walker = func(node *tree.Tree, ch chan int) {
// 		if node == nil {
// 			return
// 		}
// 		walker(node.Left, ch)
// 		ch <- node.Value
// 		walker(node.Right, ch)
// 	}

// 	walker(t, ch)
// 	close(ch)
// }

// // Same determines whether the trees
// // t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)
// 	go Walk(t1, ch1)
// 	go Walk(t2, ch2)
// 	for {
// 		v1, ok1 := <-ch1
// 		v2, ok2 := <-ch2
// 		if ok1 != ok2 && v1 != v2 {
// 			return false
// 		}

// 		if !ok1 {
// 			break
// 		}
// 	}
// 	return true
// }

// func main() {
// 	ch := make(chan int)
// 	go Walk(tree.New(1), ch)
// 	for v, ok := <-ch; ok != false; v, ok = <-ch {
// 		fmt.Printf("%v ", v)
// 	}
// 	sameTrees := Same(tree.New(1), tree.New(1))
// 	fmt.Printf("\n%v", sameTrees)
// }
