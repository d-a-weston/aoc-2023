package main

import (
	"testing"
)

func TestColorMax(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{"Color max", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", map[string]int{"blue": 6, "red": 4, "green": 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := colorMax(test.input)

			for color, value := range test.expected {
				if actual[color] != value {
					t.Errorf("Expected %d, got %d", value, actual[color])
				}
			}
		})
	}
}
