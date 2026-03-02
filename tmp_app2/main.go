package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	adminUser, err := user.NewAdmin("admin@example.com", "password")
	if err != nil {
		fmt.Println("Error creating admin:", err)
		return
	}

	//outputUserDetails(&appUser)
	appUser.OutputUserDetails()
	appUser.ClearuserName()
	appUser.OutputUserDetails()

	// adminUser.User.OutputUserDetails()
	// adminUser.User.ClearuserName()
	// adminUser.User.OutputUserDetails()
	adminUser.OutputUserDetails()
	adminUser.ClearuserName()
	adminUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
