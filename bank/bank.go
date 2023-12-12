package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) { 
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err  = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		panic("Can't continue, sorry.")
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("How can we help you today?")
		fmt.Println("1. Check Account Balance")
		fmt.Println("2. Make a Deposit")
		fmt.Println("3. Withdraw Funds")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Please choose an option (1-4): ")
		fmt.Scan(&choice)


		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("How much do you want to deposit? ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid input for deposit amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Deposit successful. Your new balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("How much do you want to widthdaw?")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid input for deposit amount. Must be greater than 0.")
				continue
			} else if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds. Cannot withdraw. ")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Widthdaw successful. Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Exiting. Thank you for using Go Bank!")
			fmt.Println("Have a great day!")
			return
		} 
	}
}
