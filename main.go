package main

import (
	"fmt"

	"example.com/tmp/utils"
)

func main() {
	fmt.Println("Hello, World!")
	testFunc()
	utils.UtilTestFunc()

	var name str = "Super Max"
	name.log()

}

type str string

func (text str) log() {
	fmt.Println(text)
}
