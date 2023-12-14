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

	board := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		board = append(board, strings.Split(line, ""))
	}

	tiltedBoard := tiltBoard(board, "north")
	total := 0

	for i, row := range tiltedBoard {
		numStones := 0

		for _, col := range row {
			if col == "O" {
				numStones++
			}
		}

		total += numStones * (len(tiltedBoard) - i)

		fmt.Println(row)
	}

	fmt.Println(total)

	file.Close()
}

func tiltBoard(board [][]string, direction string) [][]string {
	tiltedBoard := [][]string{}

	for i := 0; i < len(board); i++ {
		tiltedBoard = append(tiltedBoard, []string{})
	}

	if direction == "north" {
		for i := 0; i < len(board[0]); i++ {
			tiltedSlice := tiltSlice(getColumnSlice(board, i))

			for j := 0; j < len(tiltedSlice); j++ {
				tiltedBoard[j] = append(tiltedBoard[j], tiltedSlice[j])
			}
		}
	}

	return tiltedBoard
}

func tiltSlice(slice []string) []string {
	i := 0
	j := 0

	for j < len(slice) {
		if slice[j] == "O" {
			slice[i] = "O"

			if i != j {
				slice[j] = "."
			}

			i++
		} else if slice[j] == "#" {
			i = j + 1
		}

		j++
	}

	return slice
}

func getColumnSlice(pattern [][]string, col int) []string {
	columnSlice := []string{}

	for _, row := range pattern {
		columnSlice = append(columnSlice, row[col])
	}

	return columnSlice
}
