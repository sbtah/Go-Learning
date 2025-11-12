package main

import "fmt"

func main() {
	cards := []string{
		"Ace Of Diamonds", newCard(),
	}
	// append function does not modify the existing slice.
	// It creates a new an reassign it to `cards` variable.
	cards = append(cards, "Six Of Spades")

	// Iterate over slice:
	for idx, card := range cards {
		fmt.Println(idx, card)
	}

	fmt.Println(cards)
}

func newCard() string {
	return "Five Of Diamonds"
}
