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

	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
