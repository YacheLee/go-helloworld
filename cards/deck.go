package main

import "fmt"

type deck []string

func newDeck() deck {
	decks := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Card"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			decks = append(decks, value+" of "+suit)
		}
	}

	return decks
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
