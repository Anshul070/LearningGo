// // Exercise: Loops and Functions
// package main

// import (
// 	"fmt"
// )

// type ErrNegativeSqrt float64

// func (e ErrNegativeSqrt) Error() string{
// 	return fmt.Sprintln("cannot Sqrt negative number: ", float64(e))
// }

// func Sqrt(x float64) (float64, error) {
// 	if x < 0 {
// 		return x, ErrNegativeSqrt(x)
// 	}
// 	z := float64(1)
// 	for a := 0; a < 10; a++ {
// 		prev := z
// 		z -= (z*z - x) / (2 * z)
// 		diff := prev - z
// 		if diff < 0 {
// 			diff = -diff
// 		}
// 		if diff < 0.00000000000001 {
// 			return z, nil
// 		}
// 	}
// 	return z, nil
// }

// func main() {
// 	fmt.Println(Sqrt(2))
// 	fmt.Println(Sqrt(-1))
// }
