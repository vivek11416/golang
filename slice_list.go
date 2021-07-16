package main

func main() {

	cards := newDeck()

	//hand, remainCards := deal(cards, 5)
	//hand.print()
	//remainCards.print()
	//cards.print()
	//fmt.Println(card)
	//fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

}
