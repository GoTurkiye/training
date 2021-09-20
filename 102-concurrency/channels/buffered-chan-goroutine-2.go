package main

import "fmt"

func main() {
	bufferedChan := make(chan int, 5)
	done := make(chan bool)

	go func(bufChan chan int, done chan bool) {
		//when the bufChan closed, loop will be break
		for val := range bufChan {
			fmt.Println(val)
		}
		fmt.Println("channel closed")
		done <-true
	}(bufferedChan, done)

	bufferedChan <-1
	bufferedChan <-2
	bufferedChan <-3
	bufferedChan <-4
	bufferedChan <-5
	bufferedChan <-6
	bufferedChan <-7
	bufferedChan <-8
	bufferedChan <-9
	close(bufferedChan)

	//wait goroutine to finish
	<-done

	fmt.Println("main finished")
}
