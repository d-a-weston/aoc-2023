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

	for fileScanner.Scan() {
		line := fileScanner.Text()

		fmt.Println(line)
	}
}

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
	var nextMap string

	mapPrefix := strings.Split(currentMap, "-")[2]

	for _, v := range maps {
		if strings.HasPrefix(v, mapPrefix) {
			nextMap = v
		}
	}

	return nextMap
}
