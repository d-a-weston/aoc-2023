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
		{"Works with multiple entries", []string{"1 3 3", "4 9 2"}, map[string]string{"1": "3", "2": "4", "3": "5", "4": "9", "5": "10"}},
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

func TestFindNextMap(t *testing.T) {
	maps := []string{
		"seed-to-soil",
		"soil-to-fertilizer",
		"fertilizer-to-water",
	}

	tests := []struct {
		name       string
		currentMap string
		expected   string
	}{
		{"Finds next map", "seed-to-soil", "soil-to-fertilizer"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findNextMap(maps, test.currentMap)

			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}
