package main

func main() {

	//cards := newDeckFromFile("my_cards")
	//cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

	//hand, remainCards := deal(cards, 5)
	//hand.print()
	//remainCards.print()
	//cards.print()
	//fmt.Println(card)
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")

}
