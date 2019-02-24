package main

// go test -benchmem -run=^$ -bench. -count=1

import (
	"testing"
)

func BenchmarkKnn(b *testing.B) {
	for n := 1; n <= b.N; n++ {
		main()
	}
}
