package main

import "fmt"

func main() {

	/* var revenue float64
	var expenses float64
	var taxRate float64
 */
	/* fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue) */

	 revenue := getUserInput("Enter Revenue: ")

	/* fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses) */

	expenses := getUserInput("Enter Expenses: ")

/* 	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate) */

	taxRate := getUserInput("Enter Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	/* ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit */

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

}

func getUserInput(infoText string) (float64) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate / 100)
	ratio = ebt / profit
	return
}