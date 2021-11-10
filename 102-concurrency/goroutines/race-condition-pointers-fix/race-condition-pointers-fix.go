package raceConditionPointersFix

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

	// mtx := &sync.Mutex{} // This line for mutex lock method

	for i:=0; i<10000; i++ {
		go incrementWithAtomic(raceTest, wg)
		// go incrementWithLock(raceTest, wg, mtx) // This line for mutex lock method
	}

	wg.Wait()

	fmt.Println(raceTest)
}

func incrementWithAtomic(rt *RaceTest, wg *sync.WaitGroup) {
	atomic.AddInt32(&rt.Val, 1)
	wg.Done()
}

func incrementWithLock(rt *RaceTest, wg *sync.WaitGroup, mtx *sync.Mutex) {
	mtx.Lock()
	rt.Val += 1
	mtx.Unlock()
	wg.Done()
}