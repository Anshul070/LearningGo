package main

import (
	"fmt"
	"runtime"
	"os"
)

func main(){
	store := make([][]byte, 10)
	fmt.Println(os.Getenv("GOGC"))
	var mem runtime.MemStats;
	for i := range 15 {
		chunk := make([]byte, 50<<20) // Allocate 50 MB chunks
		store = append(store, chunk)
		runtime.ReadMemStats(&mem)
		if (i%3) == 0 {
			fmt.Printf("HeapAlloc : %v 	NumGC : %v\n", mem.HeapAlloc /1024/1024, mem.NumGC);
		}
	}
	fmt.Printf("HeapAlloc : %v 	NumGC : %v	TotalAlloc : %v\n", mem.HeapAlloc /1024/1024, mem.NumGC, mem.TotalAlloc /1024/1024);
}