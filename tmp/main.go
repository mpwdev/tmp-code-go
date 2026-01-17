package main

import "fmt"

func main() {
	const a float64 = 5.1
	const b = 6.7

	const c float64 = a * b
	const str = "Hello " + "Go"

	const d = 5 > 10

	fmt.Println(d)

	const x = 5
	const y = 6.7 * 2
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	var i int = x     // x to int
	var j float64 = x // var j float64 = float64(x)
	var p byte = x

	fmt.Println(i, j, p)
	fmt.Printf("i is: %T | j is: %T | p is: %T", i, j, p)
}
