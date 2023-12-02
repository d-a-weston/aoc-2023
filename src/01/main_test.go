package main

import (
	"testing"
)

func TestFindNumInLine(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected int
	}{
		{"Numbers at first and last", "1abc2", 12},
		{"Numbers not at first and last", "pqr3stu8vwx", 38},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findNumInLine(test.input)

			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}
