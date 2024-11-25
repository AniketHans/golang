package main

import "fmt"

// creating a new type of []string called deck
type deck []string

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