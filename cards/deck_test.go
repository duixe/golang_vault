package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck();

	if len(d) != 16 {
		t.Errorf("Expected length to be 16, but received : %v", len(d));
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First element is supposed to be %v, but received %v", "Spades", d[0]);
	}

	if d[len(d) -1] != "Four of Club" {
		t.Errorf("Last element is supposed to be %v, but received %v", "Club", d[len(d) -1]);
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_deckTesting");

	d := newDeck();

	d.saveToFile("_deckTesting");

	loadedDeck := d.readFromFile("_deckTesting");

	if (len(loadedDeck) != 16) {
		t.Errorf("Loaded deck is supposed to be 16, but receieved a length of %v", len(loadedDeck));
	}

	os.Remove("_deckTesting");
}