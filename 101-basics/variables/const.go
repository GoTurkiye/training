package main

const sabit_1 = "deger 1"

const (
	sabit_2 = "deger 2"
	sabit_3 = "deger 3"
	sabit_4 = "deger 4"
)

const (
	sabit_5 = iota
	sabit_6
	sabit_7
)

func main() {
	//sabit_1 = "deger 1.1"

	println(sabit_1, sabit_2, sabit_3, sabit_4, sabit_5, sabit_6, sabit_7)
}
