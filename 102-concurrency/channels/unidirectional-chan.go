package main

import "fmt"

func writeNumbers(chNumbers chan<- int) { //you can only send data Unidirectional  channel
	for i := 0; i < 10; i++ {
		chNumbers <- i
	}
	close(chNumbers)
}

func takeSquare(chNumbers <-chan int, chSquare chan<- int) {
	for chNumber := range chNumbers {
		chSquare <- chNumber * chNumber
	}
	close(chSquare)
}

func showResult(chSquare <-chan int) { //you can only read data chSquare Unidirectional channel
	for sqNumber := range chSquare {
		fmt.Println(sqNumber)
	}

}

func main() {
	//bidirectional-Channel
	numbers := make(chan int)
	squareNumbers := make(chan int)

	go writeNumbers(numbers)
	go takeSquare(numbers, squareNumbers)

	showResult(squareNumbers)

	//bidirectional-Channel can transform Unidirectional-Channel but other direction is not possible
}
