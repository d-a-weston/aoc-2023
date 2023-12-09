package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)

	directions := ""
	startingPositions := []string{}
	maps := map[string]map[string]string{}
	lineNumber := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if lineNumber == 0 {
			directions = line
		}

		re := regexp.MustCompile(`[A-Z]{3}`)
		matches := re.FindAll([]byte(line), -1)

		if len(matches) == 3 {
			maps[string(matches[0])] = map[string]string{
				"L": string(matches[1]),
				"R": string(matches[2]),
			}

			if strings.HasSuffix(string(matches[0]), "A") {
				startingPositions = append(startingPositions, string(matches[0]))
			}
		}

		lineNumber++
	}

	steps := 0
	currentStep := 0
	currentMaps := startingPositions

	for !All(currentMaps, func(m string) bool { return strings.HasSuffix(m, "Z") }) {
		steps++

		nextMaps := []string{}

		for _, m := range currentMaps {
			nextMaps = append(nextMaps, maps[m][string(directions[currentStep])])
		}

		if currentStep == len(directions)-1 {
			currentStep = 0
		} else {
			currentStep++
		}

		currentMaps = nextMaps
		fmt.Printf("%d: %v\n", steps, nextMaps)
	}

	fmt.Println(steps)

	file.Close()
}

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}
