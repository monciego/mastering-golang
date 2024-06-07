package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	messageSentCost := calculateBaseBill(costPerMessage, numMessages)
	discountRate := calculateDiscount(numMessages)
	finalBill := messageSentCost * (1 - discountRate)

	return finalBill
}

func calculateDiscount(messagesSent int) float64 {
	if messagesSent > 5000 {
		return 0.20
	} else if messagesSent > 1000 {
		return 0.10
	} else {
		return 0.00
	}
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}