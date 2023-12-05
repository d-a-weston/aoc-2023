package main

import (
	"testing"
)

func TestCreateMap(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected map[string]string
	}{
		{"Works with single entry", []string{"0 10 2"}, map[string]string{"0": "10", "1": "11"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := createMap(test.input)

			for k, v := range test.expected {
				if actual[k] != v {
					t.Errorf("Expected %s to be %s, got %s", k, v, actual[k])
				}
			}
		})
	}
}
