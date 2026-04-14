package main

import (
	"fmt"
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

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

}
