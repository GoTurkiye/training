package main

import "testing"

// BenchmarkFib1-12    	  400746	      3370 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFib1(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Fib1(15)
	}
}

// BenchmarkFib2-12    	623034940	         1.992 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFib2(b *testing.B) {
	n := 15
	memo := make([]int, n)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Fib2(n, memo)
	}
}

// BenchmarkFib3-12    	126282114	         9.108 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFib3(b *testing.B) {
	n := 15
	memo := make([]int, n + 1)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Fib3(15, memo)
	}
}