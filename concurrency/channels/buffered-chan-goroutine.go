package main

import "fmt"

func main() {
	//channel initialization
	bufferedChan := make(chan int, 5)

	go func(bufChan chan int) {
		for {
			val := <- bufChan
			fmt.Println(val)
		}
	}(bufferedChan)

	bufferedChan <-1
	bufferedChan <-2
	bufferedChan <-3
	bufferedChan <-4
	bufferedChan <-5
	bufferedChan <-6
	bufferedChan <-7
	bufferedChan <-8
	bufferedChan <-9

	//non-deterministic, not all values printing to output
	//main goroutine stop without waiting goroutine to print all values
}
