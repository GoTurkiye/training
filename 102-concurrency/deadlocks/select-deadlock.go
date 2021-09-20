package main

import "fmt"

func main() {

	chan1 := make(chan int)
	chan2 := make(chan int)

	go func(c chan int) {
		c <- 1
	}(chan1)
	go func(c chan int) {
		c <- 2
	}(chan2)

	for {
		select {
		case c1Val := <-chan1:
			fmt.Println(c1Val)
		case c2Val := <-chan2:
			fmt.Println(c2Val)
		}
	}
}
