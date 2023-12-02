package main

import (
	"fmt"
	"os"
	"strconv"
)

// https://adventofcode.com/2023/day/1
func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
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
