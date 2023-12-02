package main

import (
	"testing"
)

func TestFindNumInLine(t *testing.T) {
	line := "1abc2"
	expected := 12
	actual := findNumInLine(line)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
