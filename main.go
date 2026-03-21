package main

import (
	"fmt"

	"example.com/tmp/utils"
)

func main() {
	fmt.Println("Hello, World!")
	testFunc()
	utils.UtilTestFunc()

	fmt.Println("Welcome to the program")
	fmt.Println("--------------------------------")
	fmt.Println("1. Show all tasks")
	fmt.Println("2. Add a new task")
	fmt.Println("3. Mark a task as completed")
	fmt.Println("4. Delete a task")
	fmt.Println("0. Exit")
	fmt.Println("--------------------------------")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scan(&choice)
	// switch choice {
	// case 1:
	// 	fmt.Println("Showing all tasks")
	// case 2:
	// 	fmt.Println("Adding a new task")
	// case 3:
	// 	fmt.Println("Marking a task as completed")
	// case 4:
	// 	fmt.Println("Deleting a task")
	// case 0:
	// 	fmt.Println("Exiting the program")
	// default:
	// 	fmt.Println("Invalid choice")
	// }

	if choice == 1 {
		fmt.Println("Showing all tasks")
	} else if choice == 2 {
		fmt.Println("Adding a new task")
	} else if choice == 3 {
		fmt.Println("Marking a task as completed")
	} else if choice == 4 {
		fmt.Println("Deleting a task")
	} else if choice == 0 {
		fmt.Println("Exiting the program")
	} else {
		fmt.Println("Invalid choice")
	}
}
