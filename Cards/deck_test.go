package main

import (
	"github.com/google/go-cmp/cmp"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLen := 16

	if len(d) != expectedLen {
		t.Errorf("Expected deck length is %d, but got %d", expectedLen, len(d))
	}
}

func TestSaveDeck(t *testing.T) {
	testFile := "_decktesting"
	os.Remove(testFile)
	d := newDeck()
	d.saveToFile(testFile)
	_, err := os.Stat(testFile)
	if os.IsNotExist(err) {
		t.Errorf("%s should exists", testFile)
	}
	os.Remove(testFile)
}

func TestNewDeckFromFile(t *testing.T) {
	testFile := "_decktesting"
	os.Remove(testFile)

	d := newDeck()
	d.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)
	expectedLen := 16
	if len(loadedDeck) != expectedLen {
		t.Errorf("expected deck length is %d, but got %d", expectedLen, len(loadedDeck))
	}
	os.Remove(testFile)
}

func TestDeal(t *testing.T) {
	d := newDeck()
	expectedLen := 5
	d1, d2 := deal(d, expectedLen)
	if len(d1) != 5 {
		t.Errorf("expected deck1 length is %d, but got %d", expectedLen, len(d1))
	}
	if len(d2) != len(d)-5 {
		t.Errorf("expected deck2 length is %d, but got %d", len(d)-expectedLen, len(d2))
	}
}

func TestShuffle(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	d1 := newDeck()
	d2 := newDeck()
	d1.shuffle()
	// d2.shuffle()
	if cmp.Equal(d1, d2) {
		// don't use reflect.DeepEqual()
		// reflect.DeepEqual() compares targets by ==
		// cmp.Equal() compares targets by Equal()
		t.Errorf("Error: failed to shuffle")
	}
}
