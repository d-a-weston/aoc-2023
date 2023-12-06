package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)

	raceTimes := []int{}
	raceDistances := []int{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		lineSplit := strings.Fields(line)

		if lineSplit[0] == "Time:" {
			for _, time := range lineSplit[1:] {
				timeInt, _ := strconv.Atoi(time)
				raceTimes = append(raceTimes, timeInt)
			}
		}

		if lineSplit[0] == "Distance:" {
			for _, distance := range lineSplit[1:] {
				distanceInt, _ := strconv.Atoi(distance)
				raceDistances = append(raceDistances, distanceInt)
			}
		}
	}

	file.Close()

	total := 1

	for i := 0; i < len(raceTimes); i++ {
		total *= numValidSolutions(raceTimes[i], raceDistances[i])
	}

	fmt.Println(total)
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
