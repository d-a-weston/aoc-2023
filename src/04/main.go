package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	total := 0
	gameNumber := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		gameNumber++

		removePrefix := strings.Split(line, ":")

		numbers := strings.Split(removePrefix[1], "|")
		winningNumbers := strings.Fields(numbers[0])
		playersNumbers := strings.Fields(numbers[1])

		commonElements := findCommonElements(winningNumbers, playersNumbers)

		gamePoints := 0

		if len(commonElements) == 1 {
			gamePoints = 1
		}

		if len(commonElements) > 1 {
			gamePoints = 1

			for i := 0; i < len(commonElements)-1; i++ {
				gamePoints *= 2
			}
		}

		fmt.Printf("Game %d has %d matches, worth %d points\n", gameNumber, len(commonElements), gamePoints)

		total += gamePoints
	}

	fmt.Println(int(total))
}

func findCommonElements(firstSlice []string, secondSlice []string) []string {
	sharedElements := []string{}

	for _, v1 := range firstSlice {
		for _, v2 := range secondSlice {
			if v1 == v2 {
				sharedElements = append(sharedElements, v1)
			}
		}
	}

	return sharedElements
}
