package main

import "testing"

func BenchmarkBitwiseSwap(b *testing.B) {
	a, c := 5, 10
	for i := 0; i < b.N; i++ {
		BitwiseSwap(&a, &c)
	}
}

func BenchmarkSyntacticSwap(b *testing.B) {
	a, c := 5, 10
	for i := 0; i < b.N; i++ {
		SyntacticSwap(&a, &c)
	}
}

func BenchmarkArithmeticSwap(b *testing.B) {
	a, c := 5, 10
	for i := 0; i < b.N; i++ {
		ArithmeticSwap(&a, &c)
	}
}
