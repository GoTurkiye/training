package main

import "fmt"

var slc_1 []int

func main() {
	slc_2 := make([]int, 0)

	//slc_2[0] = 2

	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)
	slc_2 = append(slc_2, 1)

	fmt.Println(slc_1, slc_2)
	fmt.Printf("slc_1 len:%d cap:%d \n", len(slc_1), cap(slc_1))
	fmt.Printf("slc_2 len:%d cap:%d \n", len(slc_2), cap(slc_2))
}
