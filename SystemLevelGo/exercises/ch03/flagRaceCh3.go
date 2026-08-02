package main

import (
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"sync"
)
var flag int32;

func raceWorker(id int){
	finished := atomic.CompareAndSwapInt32(&flag, 0, 1)
	if finished {
		fmt.Printf("The worker who won is : %v\n", id)
	}
}

func main () {
	var wg sync.WaitGroup
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "go run <filename> <num_workers>")
		os.Exit(1)
	}
	
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Number of workers should be integer and more than 1")
		os.Exit(1)
	}

	for i := range numWorkers {
		wg.Go(func () {
			raceWorker(i+1)
		})
	}

	wg.Wait()
	fmt.Println("Program completed")
}