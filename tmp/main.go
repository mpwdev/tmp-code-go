package main

import (
	"fmt"
)

func main() {

	var x uint8 = 255
	x++ // overflow
	fmt.Println(x)

	// a := int8(255 + 1)

}
