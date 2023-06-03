package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create a function to create a deck of cards based on
// 1) Slice of CardSuits
// 2) Slice of CardValues
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}
