package main

import "fmt"

func main() {
	const days int = 7

	const n, m int = 4, 5
	const n1, m1 = 6, 7

	const (
		min1 = -500
		min2 = -300
		min3 = 200
	)

	const ( // group constants
		max1 = 500
		max2
		max3
	)

	fmt.Println(min1, min2, min3)
	fmt.Println(max1, max2, max3)

}
