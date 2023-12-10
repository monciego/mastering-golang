package main

import (
	"fmt"
	"math"
)

func main() {
	weight := getUserInput("Enter Weight: ")
	height := getUserInput("Enter Height: ")

	bmi := calculateBMI(weight, height)

	roundAndPrint(bmi)
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)
	return userInput
}

func calculateBMI(weight, height float64)  float64 {
	return weight / (height * height)
}

func roundAndPrint(value float64) {
	formattedBMI := fmt.Sprintf("Your BMI is %0.f", math.Round(value)) 
	fmt.Print(formattedBMI)
}
