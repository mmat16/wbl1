package main

import "testing"

var (
	first     = []int{1, 2, 3, 2, 3}
	second    = []int{5, 4, 1, 3, 1, 3, 1}
	firstStr  = []string{"h", "e", "l", "l", "o"}
	secondStr = []string{"t", "h", "e", "r", "e"}
)

func BenchmarkIntersectionON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntersectionON(first, second)
		IntersectionON(firstStr, secondStr)
	}
}

func BenchmarkIntersectionONSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntersectionONSquare(first, second)
		IntersectionONSquare(firstStr, secondStr)
	}
}
