package main

import (
	"fmt"
)

func init() {
	fmt.Println("this is init function #1")
}

func main() {

	fmt.Println("this is main function")

}

func init() {
	fmt.Println("this is init function #3")
}

func init() {
	fmt.Println("this is init function #2")
}
