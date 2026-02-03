package main

import (
	"fmt"
	"unsafe"
)

func main() {
	cars := []string{"Toyota", "Honda", "Nissan", "Ford"}
	fmt.Println(cars)

	newCars := []string{}
	newCars = append(newCars, cars[0:2]...)
	fmt.Println(newCars)

	cars[0] = "BMW"
	fmt.Println(cars)
	fmt.Println(newCars)

	s1 := []int{1, 2, 3, 4, 5}
	newSlice := s1[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = s1[2:5]
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p\n", &s1)

	fmt.Printf("%p %p %p\n", &s1, &newSlice, &newSlice[0])

	a := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("array's size in bytes: %d\n", unsafe.Sizeof(a))
	fmt.Printf("slice's size in bytes: %d\n", unsafe.Sizeof(s))

}
