package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Date: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Println("Future Value (Adjusted for Inflation):", futureRealValue)
	fmt.Printf("Future Value: %v\nFuture Value (Adjusted for Inflation): %v", futureValue, futureRealValue)
}
