package main

import "testing"

const s = "моя гордыня это дыня гор а не какой-нибудь равнинный помидор"

func BenchmarkReverseWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseWords(s)
	}
}

func BenchmarkReverseWordsConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseWordsConcat(s)
	}
}
