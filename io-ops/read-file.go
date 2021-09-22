package main

import (
	"fmt"
	"os"
)

func main() {
	readTest, err := os.ReadFile("testFile.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(readTest))
	}

	readedBytes, err := os.ReadFile("testFile.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(readedBytes))
	}
}