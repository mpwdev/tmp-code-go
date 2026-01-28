package main

import (
	"fmt"
)

func main() {
	grades := [4]int{
		1: 10,
		0: 5,
		2: 7,
		3: 9,
	}
	fmt.Println(grades)

	// array int default
	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	// array strings default
	names := [...]string{5: "Bob"}
	fmt.Println(names, len(names))
}
