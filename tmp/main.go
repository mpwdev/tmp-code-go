package main

import "fmt"

func main() {

	type second = uint

	var hour second = 3600
	fmt.Printf("Minutes in an hour: %d \n", hour/60)

}
