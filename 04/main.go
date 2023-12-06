package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	day04()
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func day04() {
	input := readFile()

	totalLines := len(input)
	// create all the cards with 1 copy
	cardCopies := make([]int, totalLines)
	for i := range cardCopies {
		cardCopies[i] = 1
	}

	// total of winning numbers
	total := 0
	for i, line := range input {
		cardTotal := 0
		winningCount := 0
		splitPrefix := strings.Split(line, ": ")
		numbers := strings.Split(splitPrefix[1], " | ")
		winning := numbers[0]
		selected := numbers[1]

		selectedMap := map[string]bool{}
		for _, s := range strings.Split(selected, " ") {
			if s != "" {
				selectedMap[s] = true
			}
		}

		for _, w := range strings.Split(winning, " ") {
			if _, ok := selectedMap[w]; ok {
				winningCount++
				if cardTotal == 0 {
					cardTotal = 1
					continue
				}
				cardTotal = cardTotal * 2
				continue
			}
		}

		total += cardTotal

		// add the winning count to the next cards as new copies
		for j := i + 1; j <= i+winningCount; j++ {
			cardCopies[j] += cardCopies[i]
		}
	}

	totalCards := 0
	// sum all the cards
	for _, a := range cardCopies {
		totalCards += a
	}

	// part one
	fmt.Println("Sum of winning numbers", total)

	// part two
	fmt.Println("Total Cards: ", totalCards)
}

func readFile() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}
