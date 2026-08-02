package main

import (
	"testing"
)


func BenchmarkPrime(b *testing.B){
	for i := 0; i < b.N ;i++{
		res := generatePrimesOptimized(50_000_000)
		_ = res
	}
}