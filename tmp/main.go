package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"
const defaultBalance = 1000.00

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		//fmt.Println("Error reading balance file, use default balance:", defaultBalance)
		return defaultBalance, errors.New("error reading balance file")
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		//fmt.Println("Error parsing balance:", err)
		return defaultBalance, errors.New("error parsing stored balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprintf("%.2f", balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	// var accountBalance = 1000.0
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("--------------------------------")
	}

	fmt.Println("Welcom to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
	fmt.Println("--------------------------------")

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
		writeBalanceToFile(accountBalance)
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
		writeBalanceToFile(accountBalance)
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
