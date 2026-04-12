package main

import (
	"fmt"

	"example.com/tmp/person"
	"example.com/tmp/utils"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
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

	fmt.Println(split(10))

}
