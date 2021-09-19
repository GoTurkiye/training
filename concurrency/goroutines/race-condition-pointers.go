package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

const incrementCount = 10000

func main() {
	raceTest := &RaceTest{}

	wg := &sync.WaitGroup{}
	wg.Add(incrementCount)

	mtx := &sync.Mutex{}
	for i := 0; i < incrementCount; i++ {
		go increment(raceTest, wg, mtx)
	}

	wg.Wait()
	fmt.Println(raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	rt.Val++
	mtx.Unlock()
	wg.Done()
}
