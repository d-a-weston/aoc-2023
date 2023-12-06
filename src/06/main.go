package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		fmt.Println(line)
	}

	file.Close()
}

func numValidSolutions(time int, distance int) int {
	totalSolutions := 0

	for i := 0; i <= time; i++ {
		velocity := i
		remainingTime := time - i
		distanceTravelled := velocity * remainingTime

		if distanceTravelled > distance {
			totalSolutions++
		}
	}

	return totalSolutions
}
