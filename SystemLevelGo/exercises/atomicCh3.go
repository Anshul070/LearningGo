package main

import (
	"fmt"
	"sync/atomic"
	"sync"
)

var counter int64
var at atomic.Int64
func incrementWorker(){
	for i := 0; i < 1000; i ++{
		atomic.AddInt64(&counter, 1)		
	}
}

func main(){
	var wg sync.WaitGroup;
	for i:= 0;i<5;i++{
		wg.Go(incrementWorker)
	}

	wg.Wait();
	fmt.Printf("Final value: %v", counter)
}