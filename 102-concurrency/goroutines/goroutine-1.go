package main

import "fmt"

func main() {
	go func() {
		fmt.Println("hello from goroutine")
	}()

	fmt.Println("hello from main")
}
