package main

import (
	"fmt"
)

func main() {
	s1 := []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3]
	fmt.Println(s1, s3, s4)

	s3[1] = 300
	fmt.Println(s1)
	fmt.Println(s3, s4)

	arr1 := [4]int{1, 2, 3, 4}
	slice1, slice2 := arr1[0:2], arr1[1:3]
	fmt.Println(arr1, slice1, slice2)

	arr1[1] = 20
	fmt.Println(arr1, slice1, slice2)

}
