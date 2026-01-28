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

	var numbers [3]int
	fmt.Println("Default array:", numbers)

	// Initialize elements by index
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println("Initialized array:", numbers)

	// Declare and initialize in one line (shorthand)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Shorthand initialization:", primes)

	// Let the compiler infer the length using "..."
	cities := [...]string{"London", "Paris", "Tokyo"}
	fmt.Println("Inferred length array:", cities)
}
