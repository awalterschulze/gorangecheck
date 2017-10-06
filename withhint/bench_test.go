package main

import (
	"math/rand"
	"testing"
)

func BenchmarkEncodeFixed64(b *testing.B) {
	buf := make([]byte, 8)
	u := rand.Uint64()
	for i := 0; i < b.N; i++ {
		encodeFixed64(buf, 0, u)
	}
}
