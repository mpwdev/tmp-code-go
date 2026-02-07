package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// success := scanner.Scan()
	// if success == false {
	// 	err = scanner.Err()
	// 	if err == nil {
	// 		log.Println("Scan completed and no error occurred")
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }
	// fmt.Println("First line of file:", scanner.Text())

	// scanner.Scan()
	// fmt.Println("Second line of file:", scanner.Text())

	// scanner.Scan()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
