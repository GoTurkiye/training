package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "lorem ipsum dolor sit amet"

	str_1 := str[:5]
	str_2 := str[len(str)-4:]
	str_3 := fmt.Sprintf("-> %s ipsum dolor sit %s", str_1, str_2)

	if strings.EqualFold(str_1, "LOrEM") {
		fmt.Println("str_1 equal to LOrEM")
	}

	if strings.HasPrefix(str, "lorem") {
		fmt.Println("lorem")
	}

	fmt.Println(str)
	fmt.Println(str_1)
	fmt.Println(str_2)
	fmt.Println(str_3)
	fmt.Println(strings.ToUpper(str))
}
