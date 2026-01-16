package main

import "fmt"

func main() {
	fmt.Println("Hello, we are comments")
	// one line comment

	/*
		comment block with multi lines
		comment line 1
		comment line 2
		comment line 3
	*/

	var maxValue = 100
	var writeToDB = true

	_, _ = maxValue, writeToDB
}
