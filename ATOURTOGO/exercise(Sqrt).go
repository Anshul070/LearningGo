// // Exercise: Loops and Functions
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func Sqrt(x float64) float64 {
// 	z := float64(1)
// 	for a := 0; a < 10; a++ {
// 		prev := z
// 		z -= (z*z - x) / (2 * z)
// 		diff := prev - z
// 		if diff < 0 {
// 			diff = -diff
// 		}
// 		if diff < 0.00000000000001 {
// 			return z
// 		}
// 	}
// 	return z
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// 	fmt.Println(math.Sqrt(2))
// }
