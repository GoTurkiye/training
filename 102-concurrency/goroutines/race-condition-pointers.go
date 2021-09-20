package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

var lock sync.Mutex

func main() {
	raceTest := &RaceTest{}

	wg := &sync.WaitGroup{}
	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		lock.Lock()
		go incrementFixed(raceTest, wg)
		//go increment(raceTest, wg)
	}

	wg.Wait()

	fmt.Println(raceTest)
}

func increment(rt *RaceTest, wg *sync.WaitGroup) {
	rt.Val += 1
	wg.Done()
}

func incrementFixed(rt *RaceTest, wg *sync.WaitGroup) {
	rt.Val += 1
	lock.Unlock()
	wg.Done()
}
