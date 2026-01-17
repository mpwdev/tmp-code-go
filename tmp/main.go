package main

import (
	"fmt"
)

func main() {
	// var r rune = 'f'
	// fmt.Printf("%T\n", r)
	// fmt.Println(r)

	// array
	var numbers = [4]int{3, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	// slice
	var names = []string{"Bob", "Kate", "Dan"}
	fmt.Printf("%T\n", names)

	// map
	balances := map[string]float64{
		"USD": 2349.5,
		"EUR": 6345.7,
	}
	fmt.Printf("%T\n", balances)

	// struct

	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	// pointer
	var zz int = 5
	ptr := &zz
	fmt.Printf("%T | %v\n", ptr, ptr)

	// function
	fmt.Printf("%T\n", f)

}

func f() {}
