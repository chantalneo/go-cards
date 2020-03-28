package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // Instructor will also address the asterisk at the core of next topic
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Diamonds" {
		t.Errorf("Expected last card to be King of Diamonds, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

// Food for Thoughts
// 1. Sample edge case:
// Create a deck
//       v
// Save to file
//       v
// File created!
//       v
// Attempt to load file
//       v
// Crash!
//
// Imagine that scenario above, let's say we write a test that creates a deck and saves it to a file, but unfortunately crashes when we attempt to load the file
// With that, our test crashes out entirely, resulting in remnant of that test file still laying around...
// So we always need to make sure that whenever we are writing test with Go we take care of any cleanup that needs to be done very manually because there is absolutely
// no handling by the Go testing framework. The Go testing framework is not going to detect that you wrote some file and automatically attempt to clean it up for you
// So in our test case, the first thing we're going to do is look into our current working directory and delete any files in there with the name of _decktesting, on top of clearing it upon a succession
