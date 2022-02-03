package main

import "fmt"

func main() {
	numbers := make([]int, 3, 5)
	printSlice(numbers)

	var nilNumbers []int
	printSlice(nilNumbers)

	if nilNumbers == nil {
		fmt.Printf("slice is nil\n")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
