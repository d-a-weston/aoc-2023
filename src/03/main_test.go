package main

import (
	"testing"
)

func TestCheckSurroundingCells(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		x        int
		y        int
		expected bool
	}{
		{"Find single digit with symbol", [][]string{{".", ".", "."}, {".", "1", "."}, {"*", ".", "."}}, 1, 1, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := checkSurroundingCells(test.input, test.x, test.y)

			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
