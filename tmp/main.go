package main

import "fmt"

func main() {
	title1, author1, year1 := "The Great Gatsby", "F. Scott Fitzgerald", 1925
	title2, author2, year2 := "1984", "George Orwell", 1949

	fmt.Println(title1, author1, year1)
	fmt.Println(title2, author2, year2)
	fmt.Println("--------------------------------")

	type book struct {
		title  string
		author string
		year   int
	}

	// type book01 struct {
	// 	title, author string
	// 	year          int
	// }

	book1 := book{title: title1, author: author1, year: year1}
	book2 := book{title: title2, author: author2, year: year2}

	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Printf("Book 1: %+v\n", book1)
	fmt.Printf("Book 2: %+v\n", book2)
}
