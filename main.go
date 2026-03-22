package main

import (
	"fmt"
	"time"

	"example.com/tmp/utils"
)

func main() {
	fmt.Println("Hello, World!")
	testFunc()
	utils.UtilTestFunc()

	p1 := Person{
		Name:      "Jim",
		Age:       35,
		Email:     "jim@example.com",
		createdAt: time.Now(),
	}

	p1.SayHello()
	printPerson(&p1)
	p1.ClearName()
	printPerson(&p1)

	p2 := newPerson("John", "john@example.com", 25)
	p2.SayHello()
	printPerson(p2)
	p2.ClearName()
	printPerson(p2)
}

// Person struct

type Person struct {
	Name      string
	Email     string
	Age       int
	createdAt time.Time
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name, "and I am", p.Age, "years old and my email is", p.Email)
}

func printPerson(p *Person) {
	fmt.Println("name:", p.Name)
	fmt.Println("age:", p.Age)
	fmt.Println("email:", p.Email)
	fmt.Println("created at:", p.createdAt.Format(time.DateTime))
}

func (p *Person) ClearName() {
	p.Name = ""
}

func newPerson(name string, email string, age int) *Person {
	return &Person{
		Name:  name,
		Email: email,
		Age:   age,
	}
}
