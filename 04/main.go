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
	total := 0

	input := readFile()
	for _, line := range input {
		cardTotal := 0
		numbers := strings.Split(line, " | ")
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
				if cardTotal == 0 {
					cardTotal = 1
					continue
				}
				cardTotal = cardTotal * 2
				continue
			}
		}

		total += cardTotal
	}

	fmt.Println(total)
}

func readFile() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}
