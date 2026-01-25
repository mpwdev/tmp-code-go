package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("os.Args", os.Args)
	fmt.Println("path", os.Args[0])
	fmt.Println("1st arg", os.Args[1])
	fmt.Println("2nd arg", os.Args[2])
	fmt.Println("3d arg", os.Args[3])
	fmt.Println("number of items in as.Args", len(os.Args))

}
