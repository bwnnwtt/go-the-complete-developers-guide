package main

func main() {

	cards := newDeck()
	cards.shuffle()
	// cards := newDeckFromFile("my_cards")
	cards.print()

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
	// fmt.Print(cards.toString())
	// cards.saveToFile("my_cards")

}
