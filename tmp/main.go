package main

import (
	"fmt"
)

func main() {

	i := 0
loop: // label
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

}
