/* i am just playing around here */
package main

import "fmt"

func main() {
	getMonthlyPrice("enterprise")
}

func monthlyPriceMessage (pennies int)  string {
	message := fmt.Sprintf("You're monthly price is %v pennies", pennies)
	return message
}

func getMonthlyPrice(tier string) int {
	var pennies int
	if tier == "basic" {
		pennies = 100.00 * 100
		fmt.Println(monthlyPriceMessage(pennies))
	} else if tier == "premium" {
		pennies = 150.00 * 100
		fmt.Println(monthlyPriceMessage(pennies))
	} else if tier == "enterprise" {
		pennies = 500.00 * 100
		fmt.Println(monthlyPriceMessage(pennies))
	} else {
		pennies = 0
		fmt.Println("Invalid tier")
	}

	return pennies
}
