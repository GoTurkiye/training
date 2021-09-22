package main

import (
	"fmt"
	"os"
)

func main()  {
	err := os.WriteFile("testFile.txt", []byte("test"), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

}
