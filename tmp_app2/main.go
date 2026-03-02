package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) outputUserDetails() {
	fmt.Println("User details:")
	fmt.Println("First name: ", u.firstName)
	fmt.Println("Last name: ", u.lastName)
	fmt.Println("Birth date: ", u.birthDate)
	fmt.Println("Created at: ", u.createdAt)
}

func (u *user) clearuserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName string, lastName string, birthDate string) (*user, error) {
	// some validation rules
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	//outputUserDetails(&appUser)
	appUser.outputUserDetails()
	appUser.clearuserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
