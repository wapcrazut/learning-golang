package main

import (
	"os"
	"testing"
)

const DefaultDeckSize = 16
const TestingDeckFileName = "_decktesting"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != DefaultDeckSize {
		t.Errorf("Expected deck length of %v, but got %v", DefaultDeckSize, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first card to be Four of Clubs, but got %v", d[0])
	}
}

func TestSaveToFileANdNewDeckFromFile(t *testing.T) {
	os.Remove(TestingDeckFileName)

	deck := newDeck()
	deck.saveToFile(TestingDeckFileName)

	bs, _ := os.ReadFile(TestingDeckFileName)

	if bs == nil {
		t.Errorf("Expected %v file to be created, none found.", TestingDeckFileName)
	}

	loadedDeck := newDeckFromFile(TestingDeckFileName)

	if len(loadedDeck) != DefaultDeckSize {
		t.Errorf("Expected deck length of %v, but got %v", DefaultDeckSize, len(loadedDeck))
	}

	os.Remove(TestingDeckFileName)
}
