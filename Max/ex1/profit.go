package main

import "fmt"

func main() {
	var monthlyRevenue float64
	fmt.Print("What is your monthly revenue value?: ")
	fmt.Scan(&monthlyRevenue)
	fmt.Println(monthlyRevenue)

	var monthlyExpenses float64
	fmt.Print("What are your monthly expenses value?: ")
	fmt.Scan(&monthlyExpenses)
	fmt.Println(monthlyExpenses)

	var taxRate float64
	fmt.Print("What is tax rate?: ")
	fmt.Scan(&taxRate)
	fmt.Println(taxRate)

	// Calculate EBT
	var ebt float64 = monthlyRevenue - monthlyExpenses

	// Calculate Tax:
	var tax float64 = ebt * taxRate

	// Calculate net income (after tax, profit)
	var netIncome float64 = ebt - tax

	// Calculate ratio (ebt / profit)
	var ratio float64 = ebt / netIncome

	fmt.Print("Your EBT is: ")
	fmt.Println(ebt)

}
