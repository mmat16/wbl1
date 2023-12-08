package main

import (
	"math/rand"
	"testing"

	"golang.org/x/exp/slices"
)

func BenchmarkQuicksort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nItems := 4096
		arr := make([]int, nItems)
		for i := range arr {
			arr[i] = rand.Int()
		}
		Quicksort(arr)
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nItems := 4096
		arr := make([]int, nItems)
		for i := range arr {
			arr[i] = rand.Int()
		}
		slices.Sort(arr)
	}
}
