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

	maps := map[string]map[string]string{}
	mapNames := []string{}

	var mapName string
	var mapBlock []string

	var seeds []string

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "seeds: ") {
			seeds = strings.Fields(strings.TrimPrefix(line, "seeds: "))

			continue
		}

		if strings.Contains(line, "-") {
			if mapName != "" {
				maps[mapName] = createMap(mapBlock)

				mapName = ""
				mapBlock = []string{}
			}

			mapName = strings.Fields(line)[0]
			mapNames = append(mapNames, mapName)

			continue
		}

		if line != "" {
			mapBlock = append(mapBlock, line)
		}
	}

	fmt.Println(maps)
	fmt.Println(mapNames)
	fmt.Println(seeds)
}

/*
 * Deprecated
 *
 * The memory and compute required for this is silly, going to do something different
 */
func createMap(lines []string) map[string]string {
	newMap := map[string]string{}

	fmt.Println(lines)

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
	var nextMap string

	mapPrefix := strings.Split(currentMap, "-")[2]

	for _, v := range maps {
		if strings.HasPrefix(v, mapPrefix) {
			nextMap = v
		}
	}

	return nextMap
}
