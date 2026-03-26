package main

import (
	"fmt"

	"example.com/tmp/utils"
)

func main() {
	fmt.Println("Hello, World!")
	testFunc()
	utils.UtilTestFunc()

	// comment

	// example array
	prices := [4]float64{10.0, 20.0, 30.0, 40.0}
	fmt.Println(prices)

}
