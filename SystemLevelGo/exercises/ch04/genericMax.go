package main

import "fmt"

type Sumable interface {
	~int | ~float64 | ~uint
}

func Sum[S Sumable](num []S) S{
	var sum S= 0;
	for _, n := range num {
		sum += n
	}
	return sum
}

func main (){
	fmt.Println(Sum([]float64{32,52,3,5,4,2,54,89}))
	fmt.Println(Sum([]uint{32,52,3,5,4,2,54,89}))
	fmt.Println(Sum([]int{32,52,3,5,4,2,54,89}))
}
