package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("my_file.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)
	bytesWritten, err := bufferedWriter.Write(
		[]byte{97, 98, 99},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available: %d\n", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString("\nJust some text")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written to buffer (not file): %d\n", bytesWritten)

	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Unflushed buffer size: %d\n", unflushedBufferSize)

	bufferedWriter.Flush()

}
