package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	str1 := strings.NewReader("read string 1")
	str2 := strings.NewReader("read string 2")
	mReader := io.MultiReader(str1, str2)

	io.Copy(os.Stdout, mReader)
}
