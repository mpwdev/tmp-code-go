package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Original numbers: ", numbers)
	doubled := doubleNumbers(&numbers)
	fmt.Println("Doubled numbers: ", doubled)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, v := range *numbers {
		dNumbers = append(dNumbers, v*2)
	}
	return dNumbers
}
