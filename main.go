package main

func main() {
	// card := newDeck()
	// card.print()
	// hand, remainingDeck := deal(card, 5)
	// fmt.Println(hand, remainingDeck)

	cards := newDeck()
	// card.saveToFile("my_cards")
	cards.shuffle()
	cards.print()
}
