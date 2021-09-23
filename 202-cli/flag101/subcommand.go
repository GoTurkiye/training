package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Create a new type for a list of Strings
type stringList []string

// Implement the flag.Value interface
func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

func main() {
	subOne := flag.NewFlagSet("one", flag.ExitOnError)
	oneCream := subOne.Bool("cream", false, "Cream")
	var sugar stringList
	subOne.Var(&sugar, "sugar", "")

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
		fmt.Println("  Sugar:", sugar.String())
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
