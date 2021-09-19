package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

func main() {
	raceTest := &RaceTest{}

	wg := &sync.WaitGroup{}
	wg.Add(10000)

	mtx := &sync.Mutex{}

	for i := 0; i < 10000; i++ {
		go incrementWithLock(raceTest, wg, mtx)
	}

	wg.Wait()

	fmt.Println(raceTest)
}

func incrementWithLock(rt *RaceTest, wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	rt.Val += 1
	mtx.Unlock()
	wg.Done()
}
