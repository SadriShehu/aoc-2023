package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	day08a()
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func day08a() {
	input := readFile()

	var instructions string
	var mapStepsLeft = make(map[string]string)
	var mapStepsRight = make(map[string]string)
	stepsRegex := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

	for i, line := range input {
		if i == 0 {
			instructions = line
		}

		groups := stepsRegex.FindStringSubmatch(line)

		if len(groups) > 0 {
			mapStepsLeft[groups[1]] = groups[2]
			mapStepsRight[groups[1]] = groups[3]
		}
	}

	currentStep := "AAA"
	steps := 0
	i := 0
	for {
		if i == len(instructions) && currentStep != "ZZZ" {
			i = 0
		}

		if currentStep == "ZZZ" {
			break
		}

		step := instructions[i]

		if step == 'L' {
			currentStep = mapStepsLeft[currentStep]
			steps++
			i++

			continue
		}
		currentStep = mapStepsRight[currentStep]
		steps++
		i++
	}

	fmt.Println(steps)
}

func readFile() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}
