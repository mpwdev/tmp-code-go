package main

import (
	"fmt"

	"example.com/tmp/utils"
)

func main() {
	//fmt.Println("Hello, World!")
	//testFunc()
	utils.UtilTestFunc()

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello,", name)
}
