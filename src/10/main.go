package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	pipes := [][]string{}
	startX := 0
	startY := 0

	lineNum := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		containsStart := slices.Index(line, "S")

		if containsStart != -1 {
			startX = containsStart
			startY = lineNum
		}

		pipes = append(pipes, line)

		lineNum++
	}

	direction := "down"
	curX := startX
	curY := startY
	steps := 0
	backToStart := false

	for !backToStart {
		if direction == "down" {
			curY++
		} else if direction == "up" {
			curY--
		} else if direction == "left" {
			curX--
		} else if direction == "right" {
			curX++
		}

		fmt.Printf("curX: %d, curY: %d\n", curX, curY)

		direction = setDirection(pipes[curY][curX], direction)

		steps++

		if pipes[curY][curX] == "S" {
			backToStart = true
		}
	}

	fmt.Println(steps / 2)

	file.Close()
}

func setDirection(pipe string, currDirection string) string {
	nextDirection := ""

	if pipe == "|" {
		nextDirection = currDirection
	} else if pipe == "-" {
		nextDirection = currDirection
	} else if pipe == "F" {
		if currDirection == "up" {
			nextDirection = "right"
		} else if currDirection == "left" {
			nextDirection = "down"
		}
	} else if pipe == "7" {
		if currDirection == "up" {
			nextDirection = "left"
		} else if currDirection == "right" {
			nextDirection = "down"
		}
	} else if pipe == "L" {
		if currDirection == "down" {
			nextDirection = "right"
		} else if currDirection == "left" {
			nextDirection = "up"
		}
	} else if pipe == "J" {
		if currDirection == "down" {
			nextDirection = "left"
		} else if currDirection == "right" {
			nextDirection = "up"
		}
	}

	return nextDirection
}
