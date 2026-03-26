package main

import (
	"fmt"

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
	fmt.Println(websites["google"])
	websites["google"] = "https://www.google.com/new"
	fmt.Println(websites["google"])

}
