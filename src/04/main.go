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

func findCommonElements(firstSlice []int, secondSlice []int) []int {
	sharedElements := []int{}

	for _, v1 := range firstSlice {
		for _, v2 := range secondSlice {
			if v1 == v2 {
				sharedElements = append(sharedElements, v1)
			}
		}
	}

	return sharedElements
}
