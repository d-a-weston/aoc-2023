package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
	}
}

func checkSurroundingCells(input [][]string, x int, y int) bool {
	ignore := []string{".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= len(input) {
			continue
		}

		for j := y - 1; j <= y+1; j++ {
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
