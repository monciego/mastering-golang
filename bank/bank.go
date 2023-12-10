package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("How can we help you today?")
	fmt.Println("1. Check Account Balance")
	fmt.Println("2. Make a Deposit")
	fmt.Println("3. Withdraw Funds")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Please choose an option (1-4): ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	}
}
