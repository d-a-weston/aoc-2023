package main

import (
	"fmt"
	"os"
)

// https://adventofcode.com/2023/day/1
func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
