package main

import (
	"fmt"

	"example.com/tmp/person"
	"example.com/tmp/utils"
)

func main() {
	fmt.Println("Hello, World!")
	testFunc()
	utils.UtilTestFunc()

	// comment

	websites := map[string]string{
		"google":   "https://www.google.com",
		"facebook": "https://www.facebook.com",
		"twitter":  "https://www.twitter.com",
	}

	for key, value := range websites {
		fmt.Println(key, value)
	}

	person, err := person.New("John Doe", "john.doe@example.com", 30)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(person)

}
