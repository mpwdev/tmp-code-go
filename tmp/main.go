package main

import (
	"fmt"
	"strconv"
)

func main() {

	i, err := strconv.Atoi("45a")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	// simple if
	if i, err := strconv.Atoi("20a"); err == nil {
		fmt.Println("no error, i is:", i)
	} else {
		fmt.Println(err)
		// handle the error
	}

}
