package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	board := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()

		board = append(board, strings.Split(line, ""))
	}

	for _, row := range board {
		fmt.Println(row)
	}

	fmt.Println()

	cycles := 1000000000
	tiltedBoard := board
	cachedRows := map[string][]string{}

	for i := 0; i < cycles; i++ {
		if i%100000 == 0 {
			fmt.Println(i)
		}

		tiltedBoard = tiltBoard(tiltedBoard, "north", cachedRows)
		tiltedBoard = tiltBoard(tiltedBoard, "west", cachedRows)
		tiltedBoard = tiltBoard(tiltedBoard, "south", cachedRows)
		tiltedBoard = tiltBoard(tiltedBoard, "east", cachedRows)
	}

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

func tiltBoard(board [][]string, direction string, cachedRows map[string][]string) [][]string {
	tiltedBoard := [][]string{}

	for i := 0; i < len(board); i++ {
		tiltedBoard = append(tiltedBoard, []string{})
	}

	if direction == "north" {
		for i := 0; i < len(board[0]); i++ {
			colSlice := getColumnSlice(board, i)
			sliceKey := fmt.Sprintf("%v", colSlice)
			tiltedSlice := []string{}

			if cachedRows[sliceKey] != nil {
				tiltedSlice = cachedRows[sliceKey]
			} else {
				tiltedSlice = tiltSlice(colSlice)
				cachedRows[sliceKey] = tiltedSlice
			}

			for j := 0; j < len(tiltedSlice); j++ {
				tiltedBoard[j] = append(tiltedBoard[j], tiltedSlice[j])
			}
		}
	}

	if direction == "south" {
		for i := 0; i < len(board[0]); i++ {
			colSlice := getColumnSlice(board, i)
			sliceKey := fmt.Sprintf("%v", colSlice)
			tiltedSlice := []string{}

			if cachedRows[sliceKey] != nil {
				tiltedSlice = cachedRows[sliceKey]
			} else {
				slices.Reverse(colSlice)
				tiltedSlice = tiltSlice(colSlice)
				cachedRows[sliceKey] = tiltedSlice
			}

			for j := 0; j < len(tiltedSlice); j++ {
				tiltedBoard[len(tiltedSlice)-j-1] = append(tiltedBoard[len(tiltedSlice)-j-1], tiltedSlice[j])
			}
		}
	}

	if direction == "east" {
		for i := 0; i < len(board); i++ {
			rowSlice := board[i]
			sliceKey := fmt.Sprintf("%v", rowSlice)
			tiltedSlice := []string{}

			if cachedRows[sliceKey] != nil {
				tiltedSlice = cachedRows[sliceKey]
			} else {
				slices.Reverse(rowSlice)
				tiltedSlice = tiltSlice(rowSlice)
				slices.Reverse(tiltedSlice)
				cachedRows[sliceKey] = tiltedSlice
			}

			tiltedBoard[i] = tiltedSlice
		}
	}

	if direction == "west" {
		for i := 0; i < len(board); i++ {
			rowSlice := board[i]
			sliceKey := fmt.Sprintf("%v", rowSlice)
			tiltedSlice := []string{}

			if cachedRows[sliceKey] != nil {
				tiltedSlice = cachedRows[sliceKey]
			} else {
				tiltedSlice = tiltSlice(rowSlice)
				cachedRows[sliceKey] = tiltedSlice
			}

			tiltedBoard[i] = tiltedSlice
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
