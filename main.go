package main

import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello, World!")
	// testFunc()
	// utils.UtilTestFunc()

	// person, err := person.New("John Doe2", "john.doe2@example.com", 30)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(person)

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
