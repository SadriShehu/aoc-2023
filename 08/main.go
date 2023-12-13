package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	instructions      string
	mapStepsLeft      = make(map[string]string)
	mapStepsRight     = make(map[string]string)
	startingPositions []string
)

func main() {
	readFile()

	start := time.Now()
	day08a()
	fmt.Printf("Execution time: %s\n", time.Since(start))

	start = time.Now()
	day08b()
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func day08a() {
	currentStep := "AAA"
	terminated := false
	steps := 0
	for !terminated {
		for _, step := range instructions {
			if currentStep == "ZZZ" {
				terminated = true
				break
			}

			if step == 'L' {
				currentStep = mapStepsLeft[currentStep]
				steps++

				continue
			}

			currentStep = mapStepsRight[currentStep]
			steps++
		}
	}

	fmt.Println("Part One:", steps)
}

func day08b() {
	output := 0
	if len(startingPositions) > 0 {
		ghostSteps := []int{}
		for _, node := range startingPositions {
			steps := 0
			terminated := false
			for !terminated {
				for _, direction := range instructions {
					if node[2] == 'Z' {
						terminated = true
						break
					}

					if direction == 'L' {
						node = mapStepsLeft[node]
						steps++
						continue
					}

					node = mapStepsRight[node]
					steps++
				}
			}

			ghostSteps = append(ghostSteps, steps)
		}

		result := ghostSteps[0]
		for i := 1; i < len(ghostSteps); i++ {
			result = leastCommonMultiple(result, ghostSteps[i])
		}
		output = result
	}

	fmt.Println("Part Two:", output)
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func leastCommonMultiple(a, b int) int {
	return a * b / greatestCommonDivisor(a, b)
}

func readFile() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(file), "\n")

	stepsRegex := regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`)

	for i, line := range input {
		if i == 0 {
			instructions = line
		}

		groups := stepsRegex.FindStringSubmatch(line)

		if len(groups) > 0 {
			mapStepsLeft[groups[1]] = groups[2]
			mapStepsRight[groups[1]] = groups[3]
			if groups[1][2] == 'A' {
				startingPositions = append(startingPositions, groups[1])
			}
		}
	}
}
