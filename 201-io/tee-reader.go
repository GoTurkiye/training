package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sReader := strings.NewReader("test string")
	//duplicates the stream
	tReader := io.TeeReader(sReader, os.Stdout)

	fmt.Println("we will start")

	readedBytes, _ := io.ReadAll(tReader)

	fmt.Println("\n --- readed string ---")

	fmt.Println(string(readedBytes))
}
