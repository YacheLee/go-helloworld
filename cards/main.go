package main

func main() {
	deck := newDeck()

	hand, remainingCards := deal(deck, 5)

	hand.print()
	remainingCards.print()
}
