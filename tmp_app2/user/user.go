package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) OutputUserDetails() {
	fmt.Println("User details:")
	fmt.Println("First name: ", u.firstName)
	fmt.Println("Last name: ", u.lastName)
	fmt.Println("Birth date: ", u.birthDate)
	fmt.Println("Created at: ", u.createdAt)
}

func (u *User) ClearuserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName string, lastName string, birthDate string) (*User, error) {
	// some validation rules
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
