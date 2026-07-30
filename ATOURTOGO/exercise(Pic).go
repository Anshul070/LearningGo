// package main

// import (
// 	"fmt"
// 	"golang.org/x/tour/pic"
// )

// func Pic(dx, dy int) [][]uint8 {
// 	s := make([][]uint8, dy) // length dy
// 	fmt.Println(s, len(s), cap(s))
// 	for x := range s {
// 		s[x] = make([]uint8, dx) //every single slice has length dx
// 		for y := range s[x] {
// 			s[x][y] = uint8(x ^ y)
// 		}
// 	}
// 	return s
// }

// func main() {
// 	pic.Show(Pic)
// }
