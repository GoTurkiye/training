package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type RaceTest struct {
	Val int32
}

func main() {
	raceTest := &RaceTest{}

	wg := &sync.WaitGroup{}
	wg.Add(10000)

	for i:=0; i<10000; i++ {
		go increment(raceTest, wg)
	}

	wg.Wait()

	fmt.Println(raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup) {
	atomic.AddInt32(&rt.Val, 1)
	wg.Done()
}
