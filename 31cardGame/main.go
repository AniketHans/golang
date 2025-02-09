package main

import "fmt"

func main() {
	// fmt.Println("Welcome to card game")
	// cards := deck{"Ace of Diamonds","King of Hearts"}
	// // cards2 := []string{"Ace of Diamonds","King of Hearts"}

	// cards.printCards()
	// // cards2.printCards() //This will throw error.

	// cards.addCard()
	// cards.printCards()

	cards := newDeck()
	cards.printCards()
	fmt.Println("******************")
	hand, remainindDeck := deal(cards, 5)
	hand.printCards()
	remainindDeck.printCards()
}
