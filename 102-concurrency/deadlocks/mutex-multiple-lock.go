package main

import (
	"fmt"
	"sync"
)

func main() {
	var i = 0

	lock := sync.Mutex{}

	lock.Lock()
	//second lock causes deadlocks
	lock.Lock()
	i++
	lock.Unlock()
	fmt.Println("hello concurrency")
}
