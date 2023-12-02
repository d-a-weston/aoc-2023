package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://adventofcode.com/2023/day/2
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

func colorMax(line string) map[string]int {
	return map[string]int{"blue": 6, "red": 4, "green": 2}
}
