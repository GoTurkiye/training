package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "", "your name")
	flag.Parse()

	if len(name) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Hello %s\n", name)
}
