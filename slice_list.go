package main

func main() {

	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "six of spades")

	cards.print()
	//fmt.Println(card)

}

func newCard() string {
	return "Five of Diamonds"
}
