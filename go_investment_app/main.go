package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedReturnRate = 5.5
	var years = 10
	var futureValue = calculate(investmentAmount, expectedReturnRate, years)

	fmt.Print(futureValue)
}

func calculate(investmentAmount int,
	expectedReturnRate float64,
	years int) float64 {
	var calculation = 
	float64(investmentAmount) * 
	math.Pow(
		1 + (expectedReturnRate / 100), 
		float64(years))
	return calculation
}