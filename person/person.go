package person

import (
	"errors"
	"fmt"
	"time"
)

// Person struct

type Person struct {
	name      string
	email     string
	age       int
	createdAt time.Time
}

func (p Person) SayHello() {
	fmt.Println("Hello, my name is", p.name, "and I am", p.age, "years old and my email is", p.email)
}

func PrintPerson(p *Person) {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("email:", p.email)
	fmt.Println("created at:", p.createdAt.Format(time.DateTime))
}

func (p *Person) ClearName() {
	p.name = ""
}

func New(name string, email string, age int) (*Person, error) {
	// add validation for needed fields
	if name == "" {
		return nil, errors.New("name is required")
	}
	if email == "" {
		return nil, errors.New("email is required")
	}
	if age <= 0 {
		return nil, errors.New("age is required")
	}
	return &Person{
		name:      name,
		email:     email,
		age:       age,
		createdAt: time.Now(),
	}, nil
}

// embedded struct

type PersonAdmin struct {
	Person
	AdminLevel int
}

func NewAdmin(name string, email string, age int, adminLevel int) (*PersonAdmin, error) {
	if adminLevel < 1 || adminLevel > 10 {
		return nil, errors.New("admin level must be between 1 and 10")
	}
	return &PersonAdmin{
		Person:     Person{name: name, email: email, age: age},
		AdminLevel: adminLevel,
	}, nil
}
