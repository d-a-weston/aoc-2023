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

	starfield := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		starfield = append(starfield, strings.Split(line, ""))
	}

	file.Close()

	// Expand Starfield
	for y := 0; y < len(starfield); y++ {
		if !slices.Contains(starfield[y], "#") {
			starfield = slices.Insert(starfield, y, starfield[y])
			y++
		}
	}

	vertContainsStar := false

	for x := 0; x < len(starfield[0]); x++ {
		for y := 0; y < len(starfield); y++ {
			if starfield[y][x] == "#" {
				vertContainsStar = true
			}
		}

		if !vertContainsStar {
			for y := 0; y < len(starfield); y++ {
				starfield[y] = slices.Insert(starfield[y], x, ".")
			}

			x++
		}

		vertContainsStar = false
	}

	planets := []struct{ x, y int }{}

	for y := 0; y < len(starfield); y++ {
		for x := 0; x < len(starfield[y]); x++ {
			if starfield[y][x] == "#" {
				planets = append(planets, struct{ x, y int }{x, y})
			}
		}
	}

	planetPairs := [][]struct{ x, y int }{}

	for i := 0; i < len(planets); i++ {
		for j := i + 1; j < len(planets); j++ {
			planetPairs = append(planetPairs, []struct{ x, y int }{planets[i], planets[j]})
		}
	}

	totalSteps := 0

	for _, pair := range planetPairs {
		steps := 0

		for pair[0].x != pair[1].x || pair[0].y != pair[1].y {
			if pair[0].x < pair[1].x {
				pair[0].x++

				steps++
			}

			if pair[0].x > pair[1].x {
				pair[0].x--

				steps++
			}

			if pair[0].y < pair[1].y {
				pair[0].y++

				steps++
			}

			if pair[0].y > pair[1].y {
				pair[0].y--

				steps++
			}
		}

		totalSteps += steps
	}

	fmt.Println(totalSteps)
}
