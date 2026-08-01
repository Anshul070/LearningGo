package main

import (
	"fmt"
	"time"
)

func sequenceGenerator(num int, number chan <- int){
	defer close(number)
	for i:= range num {
		number <- i
		time.Sleep(100 * time.Millisecond)
	}
}

func sequenceSquare(number <- chan int, square chan <- int){
	defer close(square)
	for i:= range number {
		square <- i*i
		time.Sleep(100 * time.Millisecond)
	}
}

func printSquence(square <- chan int, done chan <- struct{}){
	for i := range square{
		fmt.Printf("Sequence Square : %v\n", i)
	}
	done <- struct{}{}
}

func main(){
	number := make(chan int)
	square := make(chan int)
	done := make(chan struct{})

	go sequenceGenerator(15, number)
	go sequenceSquare(number, square)
	go printSquence(square, done)

	<- done
	fmt.Printf("\nProgram Completed\n")
}