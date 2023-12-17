package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	mirrorGrid := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		mirrorGrid = append(mirrorGrid, strings.Split(line, ""))

		fmt.Println(line)
	}

	energisedTiles := map[string]int{}

	dirs := map[string][]int{
		"up":    {0, -1},
		"down":  {0, 1},
		"right": {1, 0},
		"left":  {-1, 0},
	}

	type Beam struct {
		x   int
		y   int
		dir string
	}

	beams := []Beam{{
		x:   0,
		y:   0,
		dir: "right",
	}}

	for len(beams) > 0 {
		beam := beams[0]
		beams = beams[1:]

		for beam.x >= 0 && beam.x < len(mirrorGrid[0]) && beam.y >= 0 && beam.y < len(mirrorGrid) {
			tile := mirrorGrid[beam.y][beam.x]
			tileKey := fmt.Sprintf("%d,%d", beam.x, beam.y)
			energisedTiles[tileKey] += 1

			if tile == "/" {
				if beam.dir == "up" {
					beam.dir = "right"
				} else if beam.dir == "down" {
					beam.dir = "left"
				} else if beam.dir == "right" {
					beam.dir = "up"
				} else if beam.dir == "left" {
					beam.dir = "down"
				}
			} else if tile == "\\" {
				if beam.dir == "up" {
					beam.dir = "left"
				} else if beam.dir == "down" {
					beam.dir = "right"
				} else if beam.dir == "right" {
					beam.dir = "down"
				} else if beam.dir == "left" {
					beam.dir = "up"
				}
			} else if tile == "|" {
				if energisedTiles[tileKey] > 1 {
					break
				}

				if beam.dir == "left" || beam.dir == "right" {
					beam.dir = "up"
					beams = append(beams, Beam{x: beam.x, y: beam.y + 1, dir: "down"})
				}
			} else if tile == "-" {
				if energisedTiles[tileKey] > 1 {
					break
				}

				if beam.dir == "up" || beam.dir == "down" {
					beam.dir = "left"
					beams = append(beams, Beam{x: beam.x + 1, y: beam.y, dir: "right"})
				}
			}

			beam.x += dirs[beam.dir][0]
			beam.y += dirs[beam.dir][1]
		}
	}

	total := len(energisedTiles)

	fmt.Println(total)

	file.Close()
}
