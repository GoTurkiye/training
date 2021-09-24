package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	subOne := flag.NewFlagSet("one", flag.ExitOnError)
	oneCream := subOne.Bool("cream", false, "Cream")
	oneSugar := subOne.String("sugar", "", "Sugar")

	subTwo := flag.NewFlagSet("two", flag.ExitOnError)
	twoTea := subTwo.Int("tea", 0, "Tea")

	if len(os.Args) < 2 {
		fmt.Println("expected 'one' or 'two' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "one":
		subOne.Parse(os.Args[2:])
		fmt.Println("subcommand 'one'")
		fmt.Println("  Cream:", *oneCream)
		fmt.Println("  Sugar:", *oneSugar)
		fmt.Println("  tail:", subOne.Args())
	case "two":
		subTwo.Parse(os.Args[2:])
		fmt.Println("subcommand 'two'")
		fmt.Println("  tea:", *twoTea)
		fmt.Println("  tail:", subTwo.Args())
	default:
		fmt.Println("expected 'one' or 'two' subcommands")
		os.Exit(1)
	}
}
