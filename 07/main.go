package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	day7a()
}

var cardPower = map[rune]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	// 'J': 11, --> Part 1, jacks are not wild
	'Q': 12,
	'K': 13,
	'A': 14,
}

var cardHands = map[int]map[string]int{
	1: make(map[string]int),
	2: make(map[string]int),
	3: make(map[string]int),
	4: make(map[string]int),
	5: make(map[string]int),
	6: make(map[string]int),
	7: make(map[string]int),
}

var cardSets = [][]string{
	0: make([]string, 0),
	1: make([]string, 0),
	2: make([]string, 0),
	3: make([]string, 0),
	4: make([]string, 0),
	5: make([]string, 0),
	6: make([]string, 0),
}

func day7a() {
	input := readFile()

	for _, line := range input {
		hand := strings.Split(line, " ")
		bid, _ := strconv.Atoi(hand[1])
		evaluateHand(hand[0], bid)
	}

	calculateRank()
}

func evaluateHand(hand string, bid int) {
	cards := make(map[rune]int)
	for _, card := range hand {
		if _, ok := cards[card]; ok {
			cards[card]++
			continue
		}
		cards[card] = 1
	}

	highest := 0
	pairs := 0
	for i, v := range cards {
		// Part 2, jacks are wild
		if i == 'J' {
			continue
		}

		if v > highest {
			highest = v
		}

		if v%2 == 0 {
			pairs++
		}
	}

	// Part 2, jacks are wild, so we need to add them to the highest count
	if strings.Contains(hand, "J") {
		highest += strings.Count(hand, "J")
		if pairs > 0 {
			pairs -= 1
		}

		if highest == 2 && pairs == 0 {
			pairs++
		}
	}

	switch {
	case highest == 1:
		cardHands[1][hand] = bid
		cardSets[0] = append(cardSets[0], hand)
	case highest == 2 && pairs == 1:
		cardHands[2][hand] = bid
		cardSets[1] = append(cardSets[1], hand)
	case highest == 2 && pairs == 2:
		cardHands[3][hand] = bid
		cardSets[2] = append(cardSets[2], hand)
	case highest == 3 && pairs == 0:
		cardHands[4][hand] = bid
		cardSets[3] = append(cardSets[3], hand)
	case highest == 3 && pairs == 1:
		cardHands[5][hand] = bid
		cardSets[4] = append(cardSets[4], hand)
	case highest == 4:
		cardHands[6][hand] = bid
		cardSets[5] = append(cardSets[5], hand)
	case highest == 5:
		cardHands[7][hand] = bid
		cardSets[6] = append(cardSets[6], hand)
	}
}

func calculateRank() {
	rank := 1
	score := 0
	for i := 0; i < len(cardSets); i++ {
		bubbleSortAlg(cardSets[i])
		for _, hand := range cardSets[i] {
			score += cardHands[i+1][hand] * rank
			rank++
		}
	}

	fmt.Println(score)
}

func bubbleSortAlg(arr []string) {
	for j := 0; j < len(arr); j++ {
		for k := j + 1; k < len(arr); k++ {
			shouldSwap := false
			for l := 0; l < len(arr[j]); l++ {
				if cardPower[rune(arr[j][l])] > cardPower[rune(arr[k][l])] {
					shouldSwap = true
					break
				} else if cardPower[rune(arr[j][l])] < cardPower[rune(arr[k][l])] {
					break
				}
			}
			if shouldSwap {
				arr[j], arr[k] = arr[k], arr[j]
			}
		}
	}
}

func readFile() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}
