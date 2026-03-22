package main

import (
	"fmt"

	"example.com/tmp/person"
	"example.com/tmp/utils"
)

func main() {
	fmt.Println("Hello, World!")
	testFunc()
	utils.UtilTestFunc()

	// p1 := person.Person{
	// 	Name:      "Jim",
	// 	Age:       35,
	// 	Email:     "jim@example.com",
	// 	CreatedAt: time.Now(),
	// }

	p1, err := person.New("Jim", "jim@example.com", 35)
	if err != nil {
		fmt.Println("Error creating person:", err)
		return
	}

	p1.SayHello()
	person.PrintPerson(p1)
	p1.ClearName()
	person.PrintPerson(p1)

	fmt.Println("--------------------------------")

	a1, err := person.NewAdmin("John", "john@example.com", 27, 5)
	if err != nil {
		fmt.Println("Error creating admin:", err)
		return
	}
	a1.SayHello()
	person.PrintPerson(&a1.Person)
	fmt.Println("Admin level:", a1.AdminLevel)

}
