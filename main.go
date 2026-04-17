package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

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
	i, j := 10, 200

	p := &i
	fmt.Println("p", p)
	fmt.Println("*p", *p)
	*p = 30
	fmt.Println("p", p)
	fmt.Println("*p", *p)
	fmt.Println("i", i)

	d := &j
	fmt.Println("d", d)
	fmt.Println("*d", *d)
	*d = *d + 50
	fmt.Println("*d", *d)
	fmt.Println("j", j)

}
