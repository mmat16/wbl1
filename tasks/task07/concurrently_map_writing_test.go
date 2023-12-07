package main

import (
	"sync"
	"testing"
)

func BenchmarkWriteWithMutex(b *testing.B) {
	m := make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go WriteWithMutex(&mu, &wg, m, i+1)
	}
	wg.Wait()
}

func BenchmarkWriteWithSyncMap(b *testing.B) {
	m := &sync.Map{}
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go WriteWithSyncMap(m, &wg, i+1)
	}
	wg.Wait()
}
