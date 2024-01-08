package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected 16 cards, got %d", len(d))
	}
}
func CheckFiles(t *testing.T) {
	// check the files in the directory

	files, err := os.ReadDir("./")
	if err != nil {
		t.Errorf("Error reading directory: %s", err)
	}
	fmt.Println(files)
}
