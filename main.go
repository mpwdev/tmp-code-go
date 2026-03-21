package main

import (
	"fmt"

	"example.com/tmp/utils"
)

func main() {
	fmt.Println("Hello, World!")
	testFunc()
	utils.UtilTestFunc()

	outputText("Hello, Master")
}

func outputText(text string) {
	fmt.Println(text)
}
