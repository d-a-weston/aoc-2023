package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/1
func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	total := 0

	for fileScanner.Scan() {
		num := findNumInLine(fileScanner.Text())
		total += num
	}

	fmt.Println(total)
}

func findNumInLine(line string) int {
	var first, last int

	for _, char := range line {
		charInt, err := strconv.Atoi(string(char))

		if err != nil {
			continue
		}

		if first == 0 {
			first = charInt
		}

		last = charInt
	}

	numString := fmt.Sprintf("%d%d", first, last)

	num, err := strconv.Atoi(numString)

	if err != nil {
		panic(err)
	}

	return num
}

func replaceTokens(line string) string {
	tokens := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	matchesRemaining := true

	for matchesRemaining {
		matchIndex := len(line)
		matchValue := ""

		for token := range tokens {
			index := strings.Index(line, token)

			if index == -1 {
				continue
			}

			if index < matchIndex {
				matchIndex = index
				matchValue = token
			}
		}

		if matchValue != "" {
			line = strings.Replace(line, matchValue, tokens[matchValue], 1)
		} else {
			matchesRemaining = false
		}
	}

	return line
}
