package main

import "fmt"

func main() {
	var monthlyRevenue float64 = getUserInput("What is your monthly revenue?: ")
	fmt.Println(monthlyRevenue)

	var monthlyExpenses float64 = getUserInput("What are your monthly expenses value?: ")
	fmt.Println(monthlyExpenses)

	var taxRate float64 = getUserInput("What is tax rate?: ")
	fmt.Println(taxRate)

	ebt, profit, ratio := calculateFinancials(monthlyRevenue, monthlyExpenses, taxRate)

	fmt.Print("Your EBT is: ")
	fmt.Printf("%.1f\n", ebt)

	fmt.Print("Your profit is: ")
	fmt.Printf("%.1f\n", profit)

	fmt.Print("Your ratio is: ")
	fmt.Printf("%.3f\n", ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
