package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	day02()
	fmt.Printf("Execution time: %s\n", time.Since(start))
}

func day02() {
	input := readFile()

	redMax := 12
	greenMax := 13
	blueMax := 14
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

		skip := false
		red, green, blue := 0, 0, 0
		for _, c := range cubes {
			cube := strings.Split(c, " ")

			skip, red = checkCubes(cube[1], "red", cube[0], redMax, red, skip)
			skip, green = checkCubes(cube[1], "green", cube[0], greenMax, green, skip)
			skip, blue = checkCubes(cube[1], "blue", cube[0], blueMax, blue, skip)
		}

		// day02b
		powersTotal += red * green * blue

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

func checkCubes(color, expectedColor, cubeNumber string, maxColor, max int, skip bool) (bool, int) {
	if color == expectedColor {
		cubeInt, _ := strconv.Atoi(cubeNumber)
		if cubeInt > maxColor {
			skip = true
		}

		if max < cubeInt {
			max = cubeInt
		}
	}

	return skip, max
}
