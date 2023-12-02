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
		{"Color max #1", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", map[string]int{"blue": 6, "red": 4, "green": 2}},
		{"Color max #2", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", map[string]int{"blue": 4, "red": 1, "green": 3}},
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
