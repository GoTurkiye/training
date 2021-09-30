package raceConditionPointersFix

import (
	"sync"
	"testing"
)

// go test -benchmem -run=^$ -bench ^BenchmarkAtomic$
func BenchmarkAtomic(b *testing.B) {
	raceTest := &RaceTest{}
	wg := &sync.WaitGroup{}
	wg.Add(b.N)

	for i := 0; i < b.N; i++ {
		go incrementWithAtomic(raceTest, wg)
	}

	wg.Wait()
}

// go test -benchmem -run=^$ -bench ^BenchmarkMutexLock$
func BenchmarkMutexLock(b *testing.B) {
	raceTest := &RaceTest{}

	wg := &sync.WaitGroup{}
	wg.Add(b.N)

	mtx := &sync.Mutex{}

	for i := 0; i < b.N; i++ {
		go incrementWithLock(raceTest, wg, mtx)
	}

	wg.Wait()
}
