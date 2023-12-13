package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	patterns := [][][]string{}
	nextPattern := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			patterns = append(patterns, nextPattern)
			nextPattern = [][]string{}

			continue
		}

		nextPattern = append(nextPattern, strings.Split(line, ""))
	}

	total := 0

	for _, pattern := range patterns {
		for _, row := range pattern {
			fmt.Println(row)
		}

		cols := checkForVerticalMirror(pattern)

		if cols != -1 {
			total += cols
		}

		rows := checkForHorizontalMirror(pattern)

		if rows != -1 {
			total += rows * 100
		}

		fmt.Println(cols, rows)

		fmt.Println()
	}

	fmt.Println(total)

	file.Close()
}

func checkForVerticalMirror(pattern [][]string) int {
	for i := 0; i < len(pattern[0])-1; i++ {
		pointer := i

		for j := i + 1; j <= len(pattern[0]); j++ {
			if pointer < 0 || j == len(pattern[0]) {
				return i + 1
			}

			if !doSlicesMatch(getColumnSlice(pattern, pointer), getColumnSlice(pattern, j)) {
				break
			}

			pointer--
		}
	}

	return -1
}

func getColumnSlice(pattern [][]string, col int) []string {
	columnSlice := []string{}

	for _, row := range pattern {
		columnSlice = append(columnSlice, row[col])
	}

	return columnSlice
}

func checkForHorizontalMirror(pattern [][]string) int {
	for i := 0; i < len(pattern)-1; i++ {
		pointer := i

		for j := i + 1; j <= len(pattern); j++ {
			if pointer < 0 || j == len(pattern) {
				return i + 1
			}

			if !doSlicesMatch(pattern[pointer], pattern[j]) {
				break
			}

			pointer--
		}
	}

	return -1
}

func doSlicesMatch(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, val := range slice1 {
		if val != slice2[i] {
			return false
		}
	}

	return true
}
