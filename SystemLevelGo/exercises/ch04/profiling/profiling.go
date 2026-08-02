package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"time"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
// generatePrimes returns all prime numbers up to n using the Sieve of Eratosthenes.
func generatePrimesOptimized(n int) []int {
	if n < 2 { return nil }
	if n == 2 { return []int{2} }

	// Only track odd numbers. Index i represents number (2*i + 1)
	// Size is roughly n/2
	size := (n - 1) / 2
	isComposite := make([]byte, size+1)

	limit := (int(math.Sqrt(float64(n))) - 1) / 2
	for i := 1; i <= limit; i++ {
		if isComposite[i] == 0 {
			p := 2*i + 1
			// Step by 2*p to only hit odd multiples
			for j := (p*p - 1) / 2; j <= size; j += p {
				isComposite[j] = 1
			}
		}
	}

	// Pre-allocate using the Prime Number Theorem estimate
	primes := make([]int, 0, int(float64(n)/math.Log(float64(n))))
	primes = append(primes, 2) // Manually add the only even prime

	for i := 1; i <= size; i++ {
		if isComposite[i] == 0 {
			primes = append(primes, 2*i+1)
		}
	}
	return primes
}


func main() {
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile);
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		defer pprof.StopCPUProfile()
		pprof.StartCPUProfile(f)
	}

	// Generate primes up to 50 Million (50 << 20 bytes equivalent threshold)
	n := 50_000_000

	fmt.Printf("Generating primes up to %d...\n", n)
	startTime := time.Now()
	
	primes := generatePrimesOptimized(n)
	
	duration := time.Since(startTime)

	fmt.Printf("Done! Found %d primes.\n", len(primes))
	fmt.Printf("Time taken: %v\n", duration)
	
	// Print the first 10 primes as a sanity check
	if len(primes) >= 10 {
		fmt.Printf("First 10 primes: %v\n", primes[:10])
	}
}
