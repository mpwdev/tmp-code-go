package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Println(b)
	fmt.Printf("b: %v, type: %T, len: %d, cap: %d\n", b, b, len(b), cap(b))

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)

	fmt.Println("--------------------------------")

	fmt.Println(s1[2:])
	fmt.Println(s1[:3])
	fmt.Println(s1[:])
	fmt.Println(s1[2:4])

	s1 = append(s1[:4], 100)
	fmt.Println(s1)

	s1 = append(s1[:4], 200)
	fmt.Println(s1)

}
