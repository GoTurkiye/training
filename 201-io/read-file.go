package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	readFile()
	//readFileBufio()
	//readFileSeek()
}

func readFile()  {
	readTest, err := os.ReadFile("testFile.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(readTest))
	}
}

func readFileBufio()  {
	readTest, err := os.Open("testFile.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		bufReader := bufio.NewReader(readTest)
		io.Copy(os.Stdout, bufReader)
	}
}

func readFileSeek() {
	f, _ := os.Open("testFile.txt")
	f.Seek(5, 1)
	readByte := make([]byte, 4)
	f.Read(readByte)

	fmt.Println(string(readByte))
}