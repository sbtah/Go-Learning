package main

import (
	"fmt"
	"math"
)

func main() {

	// Constant variable:
	const inflationRate float64 = 2.41

	// Declared Investment Amount variable:
	var investmentAmount float64
	fmt.Print("Amount you want to invest: ")
	fmt.Scan(&investmentAmount)

	// Years for investment:
	var yearsForInvestment float64
	fmt.Print("For how many years: ")
	fmt.Scan(&yearsForInvestment)

	// Return ratio:
	var expectedReturnRate float64
	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	var futureIncome float64 = calculateFutureValue(investmentAmount, expectedReturnRate, yearsForInvestment)

	var futureRealIncome float64 = calculateFutureRealValue(investmentAmount, expectedReturnRate, yearsForInvestment, inflationRate)

	// var formattedFutureIncome string = fmt.Sprintf("Future Income: %.1f\n", futureIncome)

	// var formattedFutureRealIncome string = fmt.Sprintf("Future Real Income: %.1f\n", futureRealIncome)

	// Output information:
	fmt.Println(":::Result:::")
	// fmt.Println("Future value: ", futureIncome)
	fmt.Printf(
		`
		Future Income: %.1f
		Future Real Income: %.1f
		`,
		futureIncome, futureRealIncome,
	)
	// fmt.Print(formattedFutureIncome, formattedFutureRealIncome)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(amount float64, returnRate float64, yearsOfInvestment float64) float64 {
	var powerResult float64 = math.Pow(1+returnRate/100, yearsOfInvestment)
	var futureValue float64 = amount * powerResult
	return futureValue
}

func calculateFutureRealValue(amount float64, returnRate float64, yearsOfInvestment float64, inflationRate float64) float64 {
	var futureValue float64 = calculateFutureValue(amount, returnRate, yearsOfInvestment)
	var powerResult float64 = math.Pow(1+inflationRate/100, yearsOfInvestment)
	var futureRealValue float64 = futureValue / powerResult
	return futureRealValue
}
