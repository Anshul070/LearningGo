// package main

// import "fmt"

// // fibonacci is a function that returns
// // a function that returns an int.
// func fibonacci() func(int) int {
// 	n1 := 0;
// 	n2 := 0;
// 	return func(n int) int {
// 		if n < 2 {
// 			return n;
// 		}
// 		n1 = fibonacci()(n-1);
// 		n2 = fibonacci()(n-2)
// 		return n1 + n2
// 	}
// }

// func main() {
// 	f := fibonacci()
// 	for i := 0; i < 20; i++ {
// 		fmt.Println(f(i))
// 	}
// }
