package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("flavor", "vanilla", "select shot flavor")
	numbPtr := flag.Int("quantity", 2, "quantity of shots")
	boolPtr := flag.Bool("cream", false, "decide if you want cream")

	var order string
	flag.StringVar(&order, "order", "complete", "status of order")

	flag.Parse()

	fmt.Println("flavor:", *wordPtr)
	fmt.Println("quantity:", *numbPtr)
	fmt.Println("cream:", *boolPtr)
	fmt.Println("order:", order)
	fmt.Println("tail:", flag.Args())
}
