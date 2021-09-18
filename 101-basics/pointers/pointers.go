package main

import (
	"fmt"
)

type String *string

func main() {
	var rstr String
	var pstr string

	fmt.Println(rstr)
	fmt.Println(pstr)

	// pstr = "go turkiye"
	// rstr = &pstr

	// fmt.Println(rstr)
	// fmt.Println(pstr)

	// pstr = "go turkiye egitim kampi"

	// fmt.Println(*rstr)
	// fmt.Println(pstr)

	// *rstr = "go turkiye egitim kampi 2021"

	// fmt.Println(*rstr)
	// fmt.Println(pstr)

	// replaceStr(pstr)

	// fmt.Println(*rstr)
	// fmt.Println(pstr)

	// replaceStr(rstr)

	// fmt.Println(*rstr)
	// fmt.Println(pstr)
}

// func replaceStr(s string) {
// 	s = strings.Replace(s, "go", "GO", 1)
// }

// func replaceStr(s String) {
// 	*s = strings.Replace(*s, "go", "GO", 1)
// }
