package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// var newFile *os.File
	// fmt.Printf("%T\n", newFile)
	// var err error

	// newFile, err = os.Create("a.txt")
	// if err != nil {
	// 	// fmt.Println("Error creating file:", err)
	// 	// return
	// 	//----
	// 	// fmt.Println(err)
	// 	// os.Exit(1)
	// 	//----
	// 	log.Fatal(err)
	// }

	// err = os.Truncate("a.txt", 0) // complete empty file
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// newFile.Close()

	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")
	p := fmt.Println
	p("File name:", fileInfo.Name())
	p("File size:", fileInfo.Size())
	p("File mode:", fileInfo.Mode())
	p("File mod time:", fileInfo.ModTime())
	p("Is Directory:", fileInfo.IsDir())
	p("Is Regular:", fileInfo.Mode().IsRegular())

	// oldPath := "a.txt"
	// newPath := "aaa.txt"
	// err = os.Rename(oldPath, newPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	err = os.Remove("aaa.txt")
	if err != nil {
		log.Fatal(err)
	}

}
