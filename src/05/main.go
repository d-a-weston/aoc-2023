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

	maps := map[string][][]string{}
	mapNames := []string{}

	var mapName string

	var seeds []string

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "seeds: ") {
			seedRanges := strings.Fields(strings.TrimPrefix(line, "seeds: "))

			for i := 0; i < len(seedRanges); i += 2 {
				startingSeed, _ := strconv.Atoi(seedRanges[i])
				currentSeedRange, _ := strconv.Atoi(seedRanges[i+1])

				for j := 0; j < currentSeedRange; j++ {
					nextSeed := strconv.Itoa(startingSeed + j)
					seeds = append(seeds, nextSeed)
				}
			}

			continue
		}

		if strings.Contains(line, "-") {
			mapName = strings.Fields(line)[0]
			mapNames = append(mapNames, mapName)

			continue
		}

		if line != "" {
			splitLine := strings.Fields(line)
			maps[mapName] = append(maps[mapName], splitLine)
		}
	}

	var closestSeed int

	for _, seed := range seeds {
		currentMap := "seed-to-soil"

		currentValue, _ := strconv.Atoi(seed)
		prevValue := currentValue

		for currentMap != "" {
			prevValue = currentValue
			fmt.Printf("Seed: %s, Current Value: %d, Current Map: %s\n", seed, currentValue, currentMap)

			for _, slice := range maps[currentMap] {
				sourceCategory, _ := strconv.Atoi(slice[1])
				rangeValue, _ := strconv.Atoi(slice[2])

				lowerBound := sourceCategory
				upperBound := sourceCategory + rangeValue

				if currentValue >= lowerBound && currentValue <= upperBound {
					destinationCategory, _ := strconv.Atoi(slice[0])

					sourceDiff := currentValue - sourceCategory

					currentValue = destinationCategory + sourceDiff
					currentMap = findNextMap(mapNames, currentMap)

					break
				}
			}

			if prevValue == currentValue {
				currentMap = findNextMap(mapNames, currentMap)
			}
		}

		if closestSeed == 0 || currentValue < closestSeed {
			closestSeed = currentValue
		}
	}

	fmt.Println(closestSeed)
}

/*
 * Deprecated
 *
 * The memory and compute required for this is silly, going to do something different
 */
func createMap(lines []string) map[string]string {
	newMap := map[string]string{}

	for _, line := range lines {
		values := strings.Split(line, " ")

		key, _ := strconv.Atoi(values[0])
		value, _ := strconv.Atoi(values[1])
		n, _ := strconv.Atoi(values[2])

		for i := 0; i < n; i++ {
			newMap[strconv.Itoa(key)] = strconv.Itoa(value)

			key++
			value++
		}
	}

	return newMap
}

func findNextMap(maps []string, currentMap string) string {
	nextMap := ""

	splitCurrentMap := strings.Split(currentMap, "-")

	fmt.Printf("currentMap: %s, splitCurrentMap: %v, len(splitCurrentMap): %d\n", currentMap, splitCurrentMap, len(splitCurrentMap))

	mapPrefix := splitCurrentMap[2]

	for _, v := range maps {
		if strings.HasPrefix(v, mapPrefix) {
			nextMap = v
		}
	}

	return nextMap
}
