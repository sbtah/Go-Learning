package main

import (
	"fmt"
	"math"
)

func main() {

	// Constant variable:
	const inflationRate float64 = 6.5

	// Declared variable:
	var investmentAmount float64

	var yearsForInvestment float64 = 10.0
	var expectedReturnRate float64 = 5.5

	fmt.Scan(&investmentAmount)

	var powerResult float64 = math.Pow(1+expectedReturnRate/100, yearsForInvestment)

	var futureIncome float64 = investmentAmount * powerResult

	var futureRealIncome float64 = futureIncome / math.Pow(1+inflationRate/100, yearsForInvestment)

	fmt.Println(futureIncome)
	fmt.Println(futureRealIncome)
}
