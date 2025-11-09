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

	var powerResult float64 = math.Pow(1+expectedReturnRate/100, yearsForInvestment)

	var futureIncome float64 = investmentAmount * powerResult

	var futureRealIncome float64 = futureIncome / math.Pow(1+inflationRate/100, yearsForInvestment)

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
