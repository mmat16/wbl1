package main

import (
	"math/rand"
	"testing"
)

func BenchmarkDeleteFromSliceOrdered(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		DeleteFromSliceOrdered(slice, rand.Int()%len(slice))
	}
}

func BenchmarkDeleteFromSliceUnordered(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		DeleteFromSliceUnordered(slice, rand.Int()%len(slice))
	}
}
