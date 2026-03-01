package main

import (
	"fmt"

	"example.com/calc/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"
const defaultBalance = 1000.00

func main() {
	// var accountBalance = 1000.0
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, defaultBalance)
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("--------------------------------")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 at:", randomdata.PhoneNumber())

	PresentOptions()

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)
	fmt.Println("================================")
	fmt.Println("your choice:", choice)
	fmt.Println("================================")

	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Print("Enter the amount to deposit: ")
		var depositAmount float64
		fmt.Scanln(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Invalid deposit amount! Must be greater than 0")
			return
		}

		accountBalance += depositAmount
		fmt.Println("Deposit successful! New balance is:", accountBalance)
		fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
	} else if choice == 3 {
		fmt.Print("Enter the amount to withdraw:")
		var withdrawAmount float64
		fmt.Scanln(&withdrawAmount)

		if withdrawAmount <= 0 {
			fmt.Println("Invalid withdraw amount! Must be greater than 0")
			return
		}

		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient balance!")
			return
		}

		accountBalance -= withdrawAmount
		fmt.Println("Withdraw successful! New balance is:", accountBalance)
		fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
	} else if choice == 4 {
		fmt.Println("Exiting...")
		return
	} else {
		fmt.Println("Invalid choice!")
	}

	// switch choice {
	// case 1:
	// 	fmt.Println("Enter the amount to deposit:")
	// 	var amount float64
	// 	fmt.Scanln(&amount)
	// 	fmt.Println("Deposit successful!")
	// case 2:
	// 	fmt.Println("Enter the amount to withdraw:")
	// 	var amount float64
	// 	fmt.Scanln(&amount)
	// 	fmt.Println("Withdraw successful!")
	// case 3:
	// 	fmt.Println("Your balance is:")
	// 	fmt.Println("Balance successful!")
	// case 4:
	// 	fmt.Println("Exiting...")
	// 	return
	// default:
	// 	fmt.Println("Invalid choice!")
	// }
}
