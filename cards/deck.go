package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func (d deck) printcard() {
	for i, card := range d {
		//Code to print line to screen
		fmt.Println(i, card)
	}
}
