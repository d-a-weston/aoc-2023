package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

var handTypes = map[int]string{
	1: "High Card",
	2: "One Pair",
	3: "Two Pair",
	4: "Three of a Kind",
	5: "Full House",
	6: "Four of a Kind",
	7: "Five of a Kind",
}

var cardValues = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func main() {
	file, _ := os.Open("input.txt")

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		fmt.Println(line)
	}

	file.Close()
}

func getHandType(hand string) int {
	var handValue int

	cardsInHand := map[string]int{}

	for _, card := range hand {
		cardsInHand[string(card)]++
	}

	numOfSameCards := []int{}

	for _, value := range cardsInHand {
		if value > 0 {
			numOfSameCards = append(numOfSameCards, value)
		}
	}

	if slices.Contains(numOfSameCards, 5) {
		handValue = 7
	} else if slices.Contains(numOfSameCards, 4) {
		handValue = 6
	} else if slices.Contains(numOfSameCards, 3) && slices.Contains(numOfSameCards, 2) {
		handValue = 5
	} else if slices.Contains(numOfSameCards, 3) {
		handValue = 4
	} else if slices.Contains(numOfSameCards, 2) && len(numOfSameCards) == 3 {
		handValue = 3
	} else if slices.Contains(numOfSameCards, 2) {
		handValue = 2
	} else {
		handValue = 1
	}

	return handValue
}

func doesHandWin(hand1 string, hand2 string) bool {
	handWins := false

	hand1Value := getHandType(hand1)
	hand2Value := getHandType(hand2)

	if hand1Value > hand2Value {
		handWins = true
	} else if hand1Value < hand2Value {
		handWins = false
	} else {
		for i := 0; i < len(hand1); i++ {
			if cardValues[string(hand1[i])] > cardValues[string(hand2[i])] {
				handWins = true
				break
			} else if cardValues[string(hand1[i])] < cardValues[string(hand2[i])] {
				handWins = false
				break
			}
		}
	}

	return handWins
}
