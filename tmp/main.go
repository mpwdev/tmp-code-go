package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// byteSlice := make([]byte, 5)
	// numberBytesRead, err := io.ReadFull(file, byteSlice)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Number of bytes read: %d\n", numberBytesRead)
	// fmt.Printf("Data read: %s\n", byteSlice)

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read: %s\n", string(data))
	fmt.Printf("Number of bytes read: %d\n", len(data))

}
