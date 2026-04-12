package main

import (
	"fmt"

	"example.com/tmp/person"
	"example.com/tmp/utils"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println("Hello, World!")
	testFunc()
	utils.UtilTestFunc()

	// comment

	person, err := person.New("John Doe2", "john.doe2@example.com", 30)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)

	//

	a, b := swap("hello", "world")
	fmt.Println(a, b)

}
