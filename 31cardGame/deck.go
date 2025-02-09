package main

import "fmt"

// creating a new type of []string called deck
type deck []string

func newDeck() deck{
	cardSuits := []string{"Spades","Hearts","Diamonds","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King"}

	cards := deck{}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, fmt.Sprintf("%v of %v",value,suit))
		}
	}

	return cards

}

func (d deck) printCards(){
	for i, card := range d{
		fmt.Println(i,card)
	}
}

func (d deck) addCard(){
	d = append(d, "Queen of diamonds")
	fmt.Println("*****Inside addCards()*****")
	d.printCards()
	fmt.Println("*****Leaving addCards()*****")
}

func deal(d deck, handsize int) (deck,deck) {
	return d[:handsize],d[handsize:]
}