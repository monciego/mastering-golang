package main

import (
	"fmt"
	"math"
)

func main() {
	height := 1.63
	weight := 69

	bmi :=  float64(weight) / (height * height)

	fmt.Println(math.Round(bmi))
}