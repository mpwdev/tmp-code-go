package main

import (
	"fmt"
	"os"
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

	if args := os.Args; len(args) != 2 {
		fmt.Println("one argument required")
	} else if km, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("the arg must be an integer. Error:", err)
	} else {
		fmt.Printf("%d km in miles is %v\n", km, float64(km)*0.621)
		fmt.Printf("%T\n", args[1])
	}

}
