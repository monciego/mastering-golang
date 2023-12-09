package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Date: ")
	outputText("Expected Return Date: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
/* 	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years) */

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (Adjusted for Inflation): %.1f", futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (Adjusted for Inflation):", futureRealValue)
	// fmt.Printf("Future Value: %.1f\nFuture Value (Adjusted for Inflation): %.1f", futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}


func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)
	return 
}