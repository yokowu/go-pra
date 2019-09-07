package main

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(5)
	}
}

func BenchmarkClimbStairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		climbStairs(5)
	}
}
