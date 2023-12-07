package main

import (
	"fmt"
	"math"
)

func main() {
	var height float64
	var weight float64

	fmt.Print("Enter weight: ")
	fmt.Scan(&weight)

	fmt.Print("Enter height: ")
	fmt.Scan(&height)

	bmi := weight / (height * height)

	fmt.Println(math.Round(bmi))
}
