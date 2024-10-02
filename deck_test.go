package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of the deck is 16 instead we got %v", len(d)+1)
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected fist card Ace of Spades but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card Four of clubs but got %s", d[len(d)-1])
	}
}
