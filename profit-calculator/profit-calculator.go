package main

import (
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Enter Revenue: ")

	if err != nil {
		fmt.Print("Error", err)
		panic("Cannot continue due to errors")
	}

	expenses, err := getUserInput("Enter Expenses: ")

	if err != nil {
		fmt.Print("Error", err)
		panic("Cannot continue due to errors")
	}

	taxRate, err := getUserInput("Enter Tax Rate: ")

	if err != nil {
		fmt.Print("Error", err)
		panic("Cannot continue due to errors")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	storeResultToFile(ebt, profit, ratio)
}

func storeResultToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %v", err)
	}

	if userInput <= 0 {
		return 0, fmt.Errorf("invalid input: must be a positive non-zero value")
	}

	return userInput, nil
}


func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}
