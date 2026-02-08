package main

import "fmt"

func main() {

	// type book struct {
	// 	title  string
	// 	author string
	// 	year   int
	// }

	// book1 := book{title: "The Great Gatsby"}
	// fmt.Println(book1.title)
	// book1.author = "F. Scott Fitzgerald"
	// book1.year = 1925
	// fmt.Println(book1.author)
	// fmt.Println(book1.year)
	// fmt.Printf("Book 1: %+v\n", book1)

	// book2 := book{title: "The Great Gatsby", author: "F. Scott Fitzgerald", year: 1925}
	// fmt.Println(book1 == book2)

	// book2copy := book2
	// book2copy.title = "1984"
	// fmt.Println(book2copy.title)
	// fmt.Println(book2.title)

	// diana := struct { // anonymous struct
	// 	firstname string
	// 	lastname  string
	// 	age       int
	// }{firstname: "Diana", lastname: "Prince", age: 25}
	// fmt.Println(diana.firstname)
	// fmt.Println(diana.lastname)
	// fmt.Println(diana.age)
	// fmt.Printf("Diana: %+v\n", diana)
	// fmt.Printf("Diana pointer: %p\n", &diana)
	// fmt.Printf("Diana's Age: %d\n", diana.age)

	// diana2 := diana
	// diana2.firstname = "Diana"
	// diana2.lastname = "Prince"
	// diana2.age = 25
	// fmt.Println(diana2.firstname)
	// fmt.Println(diana2.lastname)
	// fmt.Println(diana2.age)

	type Book struct {
		string
		float64
		bool
	}

	b1 := Book{
		"The Great Gatsby",
		1925,
		true,
	}
	fmt.Println(b1)
	fmt.Println(b1.string)
}
