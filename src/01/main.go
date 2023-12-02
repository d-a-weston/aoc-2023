package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	return "111"
}
