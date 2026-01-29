package main

import (
	"fmt"
)

func main() {
	var cities []string
	fmt.Println("cities is equal to nill:", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	numbers := []int{2, 3, 4, 5}
	fmt.Println(numbers)

	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	var n []int
	n = numbers
	fmt.Println(n)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	myFriend1 := friends[0]
	fmt.Println("slice 0 index string:", myFriend1)

	friends[0] = "Bob"
	fmt.Println("slice 0 index string:", friends[0])
}
