package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	input := [][]string{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		input = append(input, strings.Split(line, ""))
	}

	total := 0

	numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	start := -1
	end := -1
	number := ""

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if slices.Contains(numbers, input[i][j]) {
				if start == -1 {
					start = j
				}

				end = j

				number += input[i][j]
			}

			if (!slices.Contains(numbers, input[i][j]) && start != -1 && end != -1) || j == len(input[i])-1 {
				numInt, _ := strconv.Atoi(number)
				isValidPart := checkSurroundingCells(input, start, end, i)

				fmt.Printf("Found %v part %s at %d,%d,%d -> %d\n", isValidPart, number, i, start, end, total)

				if isValidPart {
					total += numInt
				}

				start = -1
				end = -1
				number = ""
			}
		}
	}

	fmt.Println(total)
}

func checkSurroundingCells(input [][]string, start int, end int, lineNumber int) bool {
	ignore := []string{".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := lineNumber - 1; i <= lineNumber+1; i++ {
		if i < 0 || i >= len(input) {
			continue
		}

		for j := start - 1; j <= end+1; j++ {
			if j < 0 || j >= len(input[i]) {
				continue
			}

			if slices.Contains(ignore, input[i][j]) {
				continue
			}

			return true
		}
	}

	return false
}
