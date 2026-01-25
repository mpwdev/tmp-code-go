package main

import (
	"fmt"
)

func main() {

	outer := 19 // var
	//_ = outer

	people := [5]string{"Bob", "Dan", "Rob", "Kate", "Marry"}
	friends := [2]string{"Kate", "Lisa"}

outer: // label
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("next line code after break")
	fmt.Println(outer) // no conflicts with label outer

}
