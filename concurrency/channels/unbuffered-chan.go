package main

import "fmt"

func main() {
	//channel initialization
	unbufferedChan := make(chan int)

	//channel declaration
	var unbufferedChan2 chan int

	fmt.Println(unbufferedChan)
	fmt.Println(unbufferedChan2) //nil

	//only can read from channel
	go func(unbufChan <-chan int) {
		value := <-unbufChan
		fmt.Println(value)

		//unbufChan <- 5 this is not work here
	}(unbufferedChan)

	unbufferedChan <- 1
}
