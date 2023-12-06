package main

import (
	"testing"
)

func TestNumValidSolutions(t *testing.T) {
	tests := []struct {
		name     string
		time     int
		distance int
		expected int
	}{
		{"Find number of valid solutions", 7, 9, 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := numValidSolutions(test.time, test.distance)

			if actual != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}
		})
	}
}
