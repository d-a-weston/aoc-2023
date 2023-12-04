package main

import (
	"bufio"
	"fmt"
	"os"
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
