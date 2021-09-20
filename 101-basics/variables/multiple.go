package main

import "fmt"

var (
	dilAdi     = "go"
	konum      = "turkiye"
	kisiSayisi int
)

func main() {
	a, b, c, d := 1, 2, "go", true
	fmt.Println(a, b, c, d)
	a, b, c, d = 2, 1, "go turkiye", false
	fmt.Println(a, b, c, d)

	var i, j int = 1, 2
	fmt.Println(i, j)
	i, j = 2, 1
	fmt.Println(i, j)

	fmt.Println(dilAdi, konum, kisiSayisi)
}
