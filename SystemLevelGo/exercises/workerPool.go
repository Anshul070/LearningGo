package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	id int
}

type Result struct {
	Id     Job
	Output string
}

func worker(id int, jobs <-chan Job, result chan<- Result) {
	for job := range jobs {
		fmt.Printf("Worker %v is workin on job %v\n", id, job.id)
		time.Sleep(500 * time.Millisecond)

		result <- Result{
			Id:     job,
			Output: fmt.Sprintf("Job %v - Worker %v", job, id),
		}
	}

	fmt.Printf("Worker %v is exiting\n", id)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "go run <filename> <num_workers> <num_jobs>\n")
		os.Exit(1)
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if numWorkers <= 0 || err != nil {
		fmt.Fprintf(os.Stderr, "Number of workers should be integer and more than 0\n")
		os.Exit(1)
	}

	numJobs, err := strconv.Atoi(os.Args[2])
	if numJobs <= 0 || err != nil {
		fmt.Fprintf(os.Stderr, "Number of Jobs should be integer and more than 0\n")
		os.Exit(1)
	}

	jobs := make(chan Job, numJobs)
	result := make(chan Result, numJobs)

	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i+1, jobs, result)
		}()
	}

	for job := range numJobs {
		jobs <- Job{id: job+1}
	}
	close(jobs)
	
	go func() {
		wg.Wait()
		// close(result)
	}()

	for res := range result {
		fmt.Printf("Result: %v\n", res.Output);
	}

	fmt.Println("All jobs processed!")
}
