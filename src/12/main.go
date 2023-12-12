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

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		config := parts[0]

		groups := []int{}

		for _, group := range strings.Split(parts[1], ",") {
			groupAsInt, _ := strconv.Atoi(group)
			groups = append(groups, groupAsInt)
		}

		total += numPossibleConfigurations(config, groups)
	}

	fmt.Println(total)

	file.Close()
}

func numPossibleConfigurations(config string, groups []int) int {
	if config == "" {
		if len(groups) == 0 {
			return 1
		}
		return 0
	}

	if len(groups) == 0 {
		if strings.Contains(config, "#") {
			return 0
		}
		return 1
	}

	result := 0

	if config[0] == '.' || config[0] == '?' {
		result += numPossibleConfigurations(config[1:], groups)
	}

	if config[0] == '#' || config[0] == '?' {
		if groups[0] <= len(config) && !strings.Contains(config[:groups[0]], ".") && (groups[0] == len(config) || config[groups[0]] != '#') {
			if groups[0] == len(config) {
				result += numPossibleConfigurations("", groups[1:])
			} else {
				result += numPossibleConfigurations(config[groups[0]+1:], groups[1:])
			}
		}
	}

	return result
}
