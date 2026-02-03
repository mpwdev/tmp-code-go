package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 3}
	numbers = append(numbers, 4)
	fmt.Println(numbers)

	numbers = append(numbers, 5, 6, 7)
	fmt.Println(numbers)

	n := []int{10, 20, 30}
	numbers = append(numbers, n...)
	fmt.Println(numbers)

	src := []int{100, 200, 300}
	dst := make([]int, 2)
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)
}
