package main

import "fmt"

func main() {
	var monthlyRevenue float64 = getMonthlyRevenue()
	fmt.Println(monthlyRevenue)

	var monthlyExpenses float64 = getMonthlyExpenses()
	fmt.Println(monthlyExpenses)

	var taxRate float64 = getTaxRate()
	fmt.Println(taxRate)

	// Calculate EBT:
	var ebt float64 = calculateEBT(monthlyRevenue, monthlyExpenses)

	// Calculate Tax:
	var tax float64 = calculateTax(ebt, taxRate)

	// Calculate net income (after tax, profit)
	var netIncome float64 = ebt - tax

	// Calculate ratio (ebt / profit)
	var ratio float64 = ebt / netIncome

	fmt.Print("Your EBT is: ")
	fmt.Println(ebt)

	fmt.Print("Your profit is: ")
	fmt.Println(netIncome)

	fmt.Print("Your ratio is: ")
	fmt.Println(ratio)
}

func getMonthlyRevenue() float64 {
	var monthlyRevenue float64
	fmt.Print("What is your monthly revenue?: ")
	fmt.Scan(&monthlyRevenue)
	return monthlyRevenue
}

func getMonthlyExpenses() float64 {
	var monthlyExpenses float64
	fmt.Print("What are your monthly expenses value?: ")
	fmt.Scan(&monthlyExpenses)
	return monthlyExpenses
}

func getTaxRate() float64 {
	var taxRate float64
	fmt.Print("What is tax rate?: ")
	fmt.Scan(&taxRate)
	return taxRate
}

func calculateEBT(revenue float64, expenses float64) float64 {
	var ebt float64 = revenue - expenses
	return ebt
}

func calculateTax(ebt float64, taxRate float64) float64 {
	var tax float64 = ebt - taxRate
	return tax
}
