//package main

//import "fmt"

//func main() {
//	fmt.Println("Hello World, This is vivek")
//}
//
package main

import "fmt"

func main1() {

	cards := []string{"One", "Two", "Three"}

	cards = append(cards, "Four")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}
