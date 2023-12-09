package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	day6a()
	fmt.Println("Execution time: ", time.Since(start))
}

func day6a() {
	input := readFile()

	// Part 1
	times := []int{}
	lengths := []int{}

	// Part 2
	singleTime := 0
	singleLength := 0
	for _, line := range input {
		data := strings.Split(line, ": ")
		re := regexp.MustCompile(`\d+`)
		d := re.FindAllString(data[1], -1)

		var buildNumber string
		for _, v := range d {
			buildNumber += v
			u, _ := strconv.Atoi(v)
			if strings.Contains(data[0], "Time") {
				times = append(times, u)
				continue
			}
			lengths = append(lengths, u)
		}

		if strings.Contains(data[0], "Time") {
			singleTime, _ = strconv.Atoi(buildNumber)
			continue
		}
		singleLength, _ = strconv.Atoi(buildNumber)
	}

	if len(times) != len(lengths) {
		log.Fatal("Invalid input")
	}

	// Solution A
	total := 1
	for i := 0; i < len(times); i++ {
		wins := calculateDistance(times[i], lengths[i])
		total *= wins
	}

	fmt.Println(total)

	// Solution B
	wins := calculateDistance(singleTime, singleLength)
	fmt.Println(wins)
}

func calculateDistance(time, length int) int {
	wins := []int{}
	for i := 0; i <= time; i++ {
		remaining := time - i
		distanceTraveled := i * remaining

		if distanceTraveled > length {
			wins = append(wins, distanceTraveled)
		}
	}

	return len(wins)
}

func readFile() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}
