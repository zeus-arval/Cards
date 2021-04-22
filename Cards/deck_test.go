package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Expected length is 16, but got %v", len(cards))
	}

	if cards[0] != "Ace of Spades"{
		t.Errorf("Expected value is Ace of Spades, but got %v",cards[0])
	}

	if cards[len(cards) - 1] != "Four of Clubs"{
		t.Errorf("Expected value is Four of Clubs, but got %v", cards[len(cards) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	path := "_decktesting"
	os.Remove(path)
	deck := newDeck()
	deck.saveToFile(path)

	loadedDeck := newDeckFromFile(path)
	if len(loadedDeck) != 16{
		t.Errorf("Expected length of deck is 16, but got %v", len(loadedDeck))
	}

	os.Remove(path)
}