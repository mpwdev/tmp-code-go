package main

import (
	"fmt"
)

func main() {

	lang := "golang"

	switch lang {
	case "Python":
		fmt.Println("it is Python")
	case "Go", "golang":
		fmt.Println("It is Go")
	default:
		fmt.Println("Other langs")
	}

}
