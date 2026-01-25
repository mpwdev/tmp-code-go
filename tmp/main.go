package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// fmt.Println("os.Args", os.Args)
	// fmt.Println("path", os.Args[0])
	// fmt.Println("1st arg", os.Args[1])
	// fmt.Println("2nd arg", os.Args[2])
	// fmt.Println("3d arg", os.Args[3])
	// fmt.Println("number of items in as.Args", len(os.Args))

	var result, err = strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("%T\n", os.Args[1])
	fmt.Printf("%T\n", result)
	fmt.Printf("%T\n", err)

}
