package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/2
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

		if err != nil {
			panic(err)
		}

		bagContents := map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		}

		colors := colorMax(line)

		if gameIsPossible(bagContents, colors) {
			total += gameNumber
		}
	}

	fmt.Println(total)
}

func colorMax(line string) map[string]int {
	colors := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	trimGames := strings.Split(line, ":")
	games := strings.Split(trimGames[1], ";")

	for _, game := range games {
		colorsInGame := strings.Split(game, ",")
		for _, color := range colorsInGame {
			color = strings.TrimSpace(color)
			colorCount := strings.Split(color, " ")
			colorName := colorCount[1]
			colorCountInt, err := strconv.Atoi(colorCount[0])

			if err != nil {
				panic(err)
			}

			if colorCountInt > colors[colorName] {
				colors[colorName] = colorCountInt
			}
		}
	}

	return colors
}

func gameIsPossible(bagContents map[string]int, colors map[string]int) bool {
	isPossible := true

	for color, count := range colors {
		if count > bagContents[color] {
			isPossible = false
		}
	}

	return isPossible
}
