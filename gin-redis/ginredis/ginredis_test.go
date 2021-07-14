package ginredis

import "testing"

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Set("first", i)
	}
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get("first")
	}
}
