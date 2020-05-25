package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println([]byte(cards.toString()))
}
