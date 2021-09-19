package main

import (
	"fmt"
	"sync"
)

func main() {
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(2)

	i := 0

	go func() {
		lock.Lock()
		i++
		lock.Unlock()

		wg.Done()
	}()

	go func( ) {
		lock.Lock()
		i++
		lock.Unlock()

		wg.Done()
	}()

	wg.Wait()
	fmt.Println(i)
}