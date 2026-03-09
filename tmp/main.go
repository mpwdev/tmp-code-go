package main

import "fmt"

type TransformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Original numbers: ", numbers)
	doubled := transformNumbers(&numbers, double)
	fmt.Println("Doubled numbers: ", doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println("Tripled numbers: ", tripled)
}

func transformNumbers(numbers *[]int, transform TransformFn) []int {
	dNumbers := []int{}
	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
