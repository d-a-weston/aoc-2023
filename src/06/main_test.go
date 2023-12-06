package main

import (
	"testing"
)

func TestPlaceholder(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"Placeholder", "Placeholder"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := "Placeholder"

			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}
