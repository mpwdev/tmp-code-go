package main

import (
	"fmt"
)

func main() {

	var x = 3
	var y = 3.2

	x = x * int(y)
	fmt.Println(x, y)
	fmt.Printf("%T\n", y)

	y = float64(x) * y
	fmt.Println(x, y)
	fmt.Printf("%T\n", x)

	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", int(y))

}
