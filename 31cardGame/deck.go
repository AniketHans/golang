package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

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


func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(filename string) error{
	err:= ioutil.WriteFile(filename, []byte(d.toString()),0666) // 0666 means anyone can read,write and modify the file
	return err
}

func newDeckFromFile(filename string) deck{
	bs,err:= ioutil.ReadFile(filename)
	if err!=nil{
		fmt.Println("Error:",err)
	}

	s := strings.Split(string(bs),",")
	return deck(s)
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d{
		newPosition := r.Intn(len(d)-1)
		d[i],d[newPosition] = d[newPosition],d[i]
	}
}

