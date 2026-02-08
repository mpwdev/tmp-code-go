package main

import "fmt"

func main() {

	type Contact struct {
		firstname string
		lastname  string
		email     string
		phone     int
	}

	type Employee struct {
		name        string
		salary      int
		contactInfo Contact
	}

	john := Employee{
		name:   "John Doe",
		salary: 100000,
		contactInfo: Contact{
			firstname: "John",
			lastname:  "Doe",
			email:     "john.doe@example.com",
			phone:     1234567890,
		},
	}

	fmt.Println(john)
	fmt.Println(john.name)
	fmt.Println(john.salary)
	fmt.Println(john.contactInfo.firstname)
	fmt.Println(john.contactInfo.lastname)
	fmt.Println(john.contactInfo.email)
	fmt.Println(john.contactInfo.phone)

	john.contactInfo.email = "john.doe@new.com"
	fmt.Println(john.contactInfo.email)
}
