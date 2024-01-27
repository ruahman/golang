package types

import (
	"os"
	"testing"
)

func TestTypes(t *testing.T) {
	Demo()
}

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	} else {
		t.Log("TestNewDeck passed")
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	} else {
		t.Log("TestNewDeck passed")
	}
}

func TestSaveToDeck(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	_ = d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	} else {
		t.Log("TestSaveToDeckAndNewDeckFromFile passed")
	}

	os.Remove("_decktesting")
}
