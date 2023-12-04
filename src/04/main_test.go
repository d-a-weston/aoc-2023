package main

import (
	"testing"
)

func TestNumSharedElements(t *testing.T) {
	tests := []struct {
		name        string
		firstSlice  []int
		secondSlice []int
		expected    []int
	}{
		{"Shares one element", []int{1, 2, 3, 4}, []int{4, 5, 6}, []int{4}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := findCommonElements(test.firstSlice, test.secondSlice)

			if len(actual) != len(test.expected) {
				t.Errorf("Expected %d, got %d", test.expected, actual)
			}

			for i, v := range actual {
				if v != test.expected[i] {
					t.Errorf("Expected %d, got %d", test.expected, actual)
				}
			}
		})
	}
}
