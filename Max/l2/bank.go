package main

import "fmt"

func main() {

	var accountBallance float64 = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check the ballance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice?: ")
	fmt.Scan(&choice)

	// var wantsCheckBallance bool = choice == 1

	if choice == 1 {
		fmt.Println("Your account ballance is: ", accountBallance)
	} else if choice == 2 {
		fmt.Print("How much do you want to deposit?: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBallance += depositAmount
		fmt.Println("Ballance updated!, New amount:", accountBallance)
	} else if choice == 3 {
		fmt.Print("How much do you want to withdraw?: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBallance {
			fmt.Println("You can't withdraw more than your current amount!")
		} else {
			accountBallance -= withdrawAmount
			fmt.Println("Ballance updated!, New amount:", accountBallance)
		}
	} else {
		fmt.Println("Goodbye")
	}

	// switch choice {
	// case 1:
	// 	fmt.Println("Your account ballance is: ", accountBallance)
	// case 2:
	// 	fmt.Println("How much do you want to deposit?")
	// }
}
