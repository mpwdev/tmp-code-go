package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age:", age)

	var name = "Bob"
	fmt.Println("Your Name is:", name)

	message1 := "Learning GoLang"
	fmt.Println(message1)
	message1 = "Learning GoLang!"
	fmt.Println(message1)

	car, cost := "Audi", 500
	fmt.Println(car, cost)

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	fmt.Println("i:", i, "j:", j)
	i, j = j, i // swapping variables
	fmt.Println("i:", i, "j:", j)
	//_, _ = i, j

	sum := 5 + 3.4
	fmt.Println(sum)
}
