package main

import "fmt"

func main() {

	// Long Version of Declaration
	// var card string = "Ace of Spades"

	//Short Version of Declaration
	//card := "Ace of Spades"

	//Note that := is only for NEW Variables, re-declaration doesnt need it
	card := newCard()

	//Code to print line to screen
	fmt.Println(card)

}

func newCard() int {
	return 5
}
