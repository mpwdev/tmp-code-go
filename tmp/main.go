package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("b.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("I learn GoLang")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten)

}
