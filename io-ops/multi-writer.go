package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	txtFile, _ := os.Create("txtFile.txt")
	txtFile2, _ := os.Create("txtFile2.txt")
	mWriter := io.MultiWriter(os.Stdout, txtFile, txtFile2)

	n, err := mWriter.Write([]byte("multi writer example"))

	fmt.Println(err)
	fmt.Println(n)
}
