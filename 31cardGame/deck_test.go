package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52{
		t.Errorf("Expected deck length of 52 but go %v",len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_deckTest")

	d := newDeck()
	d.saveToFile("_deckTest")
	
	loadedDeck := newDeckFromFile("_deckTest")
	if len(loadedDeck)!=52{
		t.Errorf("Expected deck length of 52 but go %v",len(loadedDeck))
	}
	
	os.Remove("_deckTest")
}