package main

import (
	"fmt"
	"os"
)

func main()  {
	//writeOsWrite()
	writeFile()
}

func writeOsWrite() {
	err := os.WriteFile("testFile.txt", []byte("test"), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

func writeFile() {
	f, err := os.Create("testFileWrite.txt")
	if err != nil {
		fmt.Println(err)
	}

	f.Write([]byte("1\n"))
	f.Write([]byte("2\n"))
	f.Write([]byte("3\n"))
	f.Write([]byte("4\n"))
	f.Write([]byte("5\n"))
	f.Close()
}