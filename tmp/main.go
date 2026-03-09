package main

import "fmt"

type TransformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	fmt.Println("Original numbers: ", numbers)
	doubled := transformNumbers(&numbers, double)
	fmt.Println("Doubled numbers: ", doubled)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println("Tripled numbers: ", tripled)

	transformerFn1 := getTransformerFn(&moreNumbers)
	transformedNumbers1 := transformNumbers(&moreNumbers, transformerFn1)
	fmt.Println("Transformed numbers 1: ", transformedNumbers1)

	transformerFn2 := getTransformerFn(&numbers)
	transformedNumbers2 := transformNumbers(&numbers, transformerFn2)
	fmt.Println("Transformed numbers 2: ", transformedNumbers2)
}

func transformNumbers(numbers *[]int, transform TransformFn) []int {
	dNumbers := []int{}
	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}
	return dNumbers
}

func getTransformerFn(numbers *[]int) TransformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
