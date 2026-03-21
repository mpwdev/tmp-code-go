package main

import (
	"fmt"

	"example.com/tmp/utils"
)

func main() {
	fmt.Println("Hello, World!")
	testFunc()
	utils.UtilTestFunc()

	var age2 int = 20

	var age *int
	//fmt.Println("age", age)
	age = &age2
	fmt.Println("age", age)
	fmt.Println("age", *age)
	//fmt.Println("age2", age2)
}
