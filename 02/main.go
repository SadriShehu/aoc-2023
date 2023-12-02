package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Color struct {
	maxAllowed int
	cubePowers int
}

func main() {
	start := time.Now()
	day02()
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func day02() {
	input := readFile()

	colors := map[string]*Color{
		"red":   {12, 0},
		"green": {13, 0},
		"blue":  {14, 0},
	}

	total := 0
	powersTotal := 0

	for _, line := range input {
		line := strings.Split(line, ": ")

		game := strings.Split(line[0], " ")
		gameID, _ := strconv.Atoi(game[1])

		rounds := strings.Split(line[1], "; ")
		cubes := []string{}
		for _, round := range rounds {
			pares := strings.Split(round, ", ")
			cubes = append(cubes, pares...)
		}

		// reset cube powers for each game
		colors["blue"].cubePowers = 0
		colors["green"].cubePowers = 0
		colors["red"].cubePowers = 0
		skip := false
		for _, c := range cubes {
			cube := strings.Split(c, " ")

			if color, ok := colors[cube[1]]; ok {
				cubeInt, _ := strconv.Atoi(cube[0])
				if cubeInt > color.maxAllowed {
					skip = true
				}

				if color.cubePowers < cubeInt {
					color.cubePowers = cubeInt
				}
			}
		}

		// day02b
		powersTotal += colors["red"].cubePowers *
			colors["green"].cubePowers *
			colors["blue"].cubePowers

		// day02a
		if !skip {
			total += gameID
		}
	}

	fmt.Println("Game Score:", total, "Max Cubes:", powersTotal)
}

func readFile() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}
