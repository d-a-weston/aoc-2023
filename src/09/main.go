package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		fmt.Println(line)
	}

	file.Close()
}

func sequenceDiff(sequence []int) []int {
	diffs := []int{}

	for i := 0; i < len(sequence)-1; i++ {
		diffs = append(diffs, sequence[i+1]-sequence[i])
	}

	return diffs
}

func findNextValue(sequence []int) int {
	if All(sequence, func(x int) bool { return x == 0 }) {
		return 0
	}

	return sequence[len(sequence)-1] + findNextValue(sequenceDiff(sequence))
}

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}