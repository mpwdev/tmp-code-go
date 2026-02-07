package main

import "fmt"

func main() {

	type book struct {
		title  string
		author string
		year   int
	}

	book1 := book{title: "The Great Gatsby"}
	fmt.Println(book1.title)
	book1.author = "F. Scott Fitzgerald"
	book1.year = 1925
	fmt.Println(book1.author)
	fmt.Println(book1.year)
	fmt.Printf("Book 1: %+v\n", book1)

	book2 := book{title: "The Great Gatsby", author: "F. Scott Fitzgerald", year: 1925}
	fmt.Println(book1 == book2)

	book2copy := book2
	book2copy.title = "1984"
	fmt.Println(book2copy.title)
	fmt.Println(book2.title)
}
