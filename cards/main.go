package main

import "fmt"

func main() {

	// Long Version of Declaration
	// var card string = "Ace of Spades"

	//Short Version of Declaration
	//card := "Ace of Spades"

	//Note that := is only for NEW Variables, re-declaration doesnt need it
	//card := newCard()

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {

		//Code to print line to screen

		fmt.Println(i, card)
	}
}

//Function return type needs to be defined, aka String
func newCard() string {
	return "Five of Diamonds"
}
