package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line
	if finalCost = bulkMessageCost; isPremiumUser {
		finalCost = bulkMessageCost - (bulkMessageCost * discountRate)
	} 
	
	if accountBalance >= finalCost {
		accountBalance = accountBalance - finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}


/* Problem 
Using the given variables, write conditional statements to calculate and update the variables.

First, set finalCost to the bulkMessageCost. If the user is premium, then the discountRate should be applied to the finalCost. A discountRate of 0.20 means the discounted price per message would be 80% of the original price.

Next, if the user has enough money in their accountBalance:

1. Process the payment by deducting the finalCost from their accountBalance
2. Print the purchaseSuccessMessage
3. Otherwise, just print the insufficientFundMessage.


*/