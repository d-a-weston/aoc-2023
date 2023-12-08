package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)

	directions := ""
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
		}

		lineNumber++
	}

	steps := 0
	currentStep := 0
	currentMap := "AAA"

	for currentMap != "ZZZ" {
		steps++

		currentMap = maps[currentMap][string(directions[currentStep])]

		if currentStep == len(directions)-1 {
			currentStep = 0
		} else {
			currentStep++
		}
	}

	fmt.Println(steps)

	file.Close()
}
