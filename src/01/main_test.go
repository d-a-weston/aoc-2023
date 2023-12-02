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
		{"More than 2 numbers", "a1b2c3d4e5f", 15},
		{"Only one number", "treb7uchet", 77},
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

func TestReplaceTokens(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{"Replace one with 1", "1one1", "111"},
		{"Replace multiple numbers", "two1nine", "219"},
		{"Replace numbers in the order they appear", "eightwothree", "8wo3"},
		{"Replace while preserving numbers", "xtwone3four", "x2ne34"},
		{"Additional AoC test case #1", "4nineeightseven2", "49872"},
		{"Additional AoC test case #2", "zoneight234", "z1ight234"},
		{"Additional AoC test case #3", "7pqrstsixteen", "7pqrst6teen"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := replaceTokens(test.input)

			if actual != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, actual)
			}
		})
	}
}
