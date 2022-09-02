package main

import (
	"testing"
)

func BenchmarkGetProducts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetProducts()
	}
}
