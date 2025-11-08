package main

import "fmt"

func main() {
	var monthlyRevenue float64
	fmt.Print("What is your monthly revenue?: ")
	fmt.Scan(&monthlyRevenue)

	fmt.Println(monthlyRevenue)

	var taxRate float64
	fmt.Print("What is tax rate?: ")
	fmt.Scan(&taxRate)

	fmt.Println(taxRate)
}
