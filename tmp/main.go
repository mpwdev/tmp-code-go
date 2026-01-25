package main

import (
	"fmt"
)

func main() {

	// for i := 0; i < 10; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	count := 0

	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Printf("%d is divisible by 13 \n", i)
			count++
		}
		if count == 5 {
			break
		}
	}

	// code after for loop
	fmt.Println("Code line (message) after for loop")

}
