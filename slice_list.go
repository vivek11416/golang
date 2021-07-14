package main

import (
	"fmt"
)

func main() {

	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "six of spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
	//fmt.Println(card)

}

func newCard() string {
	return "Five of Diamonds"
}
