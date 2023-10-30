package main

import "testing"

// go test -bench=. -test.benchmem

func BenchmarkCreateHugeStringSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = createHugeStringSlow(1 << 10)
	}
}

func BenchmarkCreateHugeStringFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = createHugeStringFast(1 << 10)
	}
}
