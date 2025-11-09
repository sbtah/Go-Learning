package main

import "fmt"

func main() {

	var yourName string

	fmt.Print("Yo there, what is your name?: ")
	fmt.Scan(&yourName)

	var yourLastName string
	fmt.Print("What is your last name?: ")
	fmt.Scan(&yourLastName)

	fmt.Println("Hello", yourName, yourLastName)
}
