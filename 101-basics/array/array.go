package main

import "fmt"

var arr_1 [3]int
var arr_2 = [5]int{1, 2, 3, 4, 5}

func main() {
	arr_3 := make([]int, 3)

	arr_3[1] = 2

	fmt.Println(arr_1, arr_2, arr_3)
	fmt.Printf("arr_1 len:%d \n", len(arr_1))
	fmt.Printf("arr_2 len:%d \n", len(arr_3))
}
