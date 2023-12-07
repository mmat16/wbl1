package main

import "testing"

func BenchmarkSetBit(b *testing.B) {
	var num int64
	for i := 0; i < b.N; i++ {
		SetBit(num, 15)
	}
}

func BenchmarkResetBit(b *testing.B) {
	var num int64
	for i := 0; i < b.N; i++ {
		ResetBit(num, 25)
	}
}

func BenchmarkIsSetBit(b *testing.B) {
	var num int64
	for i := 0; i < b.N; i++ {
		IsSetBit(num, 11)
	}
}
