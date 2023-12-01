package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	digitsMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func main() {
	start := time.Now()
	day01a()
	fmt.Println("Time:", time.Since(start))

	startb := time.Now()
	day01b()
	fmt.Println("Time:", time.Since(startb))
}

func day01a() {
	data := readFile()

	coordinates := []int{}
	for _, line := range data {
		num := []int{}
		for _, char := range line {
			number, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			num = append(num, number)
		}

		coordinates = append(coordinates, buildCoordinates(num))
	}

	fmt.Println("Total:", calculateTotal(coordinates))
}

func day01b() {
	data := readFile()

	lock := sync.Mutex{}
	coordinates := []int{}
	wg := sync.WaitGroup{}
	for _, line := range data {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			num := []int{}

			current := ""
			for _, char := range line {
				current = current + string(char)
				for _, digit := range digits {
					if strings.Contains(current, digit) {
						num = append(num, digitsMap[digit])
						current = string(char)
					}
				}

				number, err := strconv.Atoi(string(char))
				if err != nil {
					continue
				}
				num = append(num, number)
			}

			lock.Lock()
			coordinates = append(coordinates, buildCoordinates(num))
			lock.Unlock()
		}(line)
	}
	wg.Wait()

	fmt.Println("Total:", calculateTotal(coordinates))
}

func readFile() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}

func buildCoordinates(num []int) int {
	var coordinate string
	if len(num) == 1 {
		coordinate = fmt.Sprintf("%d%d", num[0], num[0])
	} else if len(num) >= 2 {
		coordinate = fmt.Sprintf("%d%d", num[0], num[len(num)-1])
	} else {
		return 0
	}

	crd, err := strconv.Atoi(coordinate)
	if err != nil {
		log.Fatal(err)
	}

	return crd
}

func calculateTotal(coordinates []int) int {
	total := 0
	for _, coordinate := range coordinates {
		total += coordinate
	}

	return total
}
