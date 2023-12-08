package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type rangeSeed struct {
	start int
	end   int
}

type seq struct {
	start int
	end   int
	diff  int
}

type set = []seq

func main() {
	startOne := time.Now()
	day5a()
	fmt.Println("Execution time: ", time.Since(startOne))

	startTwo := time.Now()
	day5b()
	fmt.Println("Execution time: ", time.Since(startTwo))
}

func getNumbers(line string) []int {
	re := regexp.MustCompile(`\d+`)
	found := re.FindAllStringIndex(line, -1)
	numbers := make([]int, 0)
	for _, match := range found {
		valueStr := line[match[0]:match[1]]
		val, _ := strconv.Atoi(valueStr)
		numbers = append(numbers, val)
	}

	return numbers
}

func sequence(line string) seq {
	numbers := getNumbers(line)

	source := numbers[1]
	dest := numbers[0]
	size := numbers[2]

	return seq{
		start: source,
		end:   source + size - 1,
		diff:  dest - source,
	}
}

func findSeed(result int, sets []set) int {
	curr := result

	for i := len(sets) - 1; i >= 0; i-- {
		var seqSelected seq
		found := false

		for j := len(sets[i]) - 1; j >= 0; j-- {
			seq := sets[i][j]
			if curr >= seq.start+seq.diff && curr <= seq.end+seq.diff {
				seqSelected = seq
				found = true
				break
			}
		}

		if found {
			curr = curr - seqSelected.diff
		}
	}

	return curr
}

func findLowestLocation(seed int, sets []set) int {
	curr := seed

	for _, set := range sets {
		var seqSelected seq
		found := false

		for _, seq := range set {
			if curr >= seq.start && curr <= seq.end {
				seqSelected = seq
				found = true
			}
		}

		if found {
			curr = curr + seqSelected.diff
		}
	}

	return curr
}

func sets(lines []string) []set {
	sets := make([]set, 0)

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			sets = append(sets, make([]seq, 0))
			continue
		}

		prevSet := len(sets) - 1
		sets[prevSet] = append(sets[prevSet], sequence(line))
	}

	return sets
}

func day5a() {
	lines := readFile()

	seeds := getNumbers(lines[0])
	sets := sets(lines)

	// Find lowest location for each seed
	solution := math.Inf(1) // positive infinity
	seedsResult := make([]int, len(seeds))
	for i, seed := range seeds {
		seedsResult[i] = findLowestLocation(seed, sets)
	}

	for _, res := range seedsResult {
		if res < int(solution) {
			solution = float64(res)
		}
	}

	fmt.Println("Solution A:", int(solution))
}

func day5b() {
	lines := readFile()

	seeds := make([]rangeSeed, 0)
	numbers := getNumbers(lines[0])
	for i := 0; i < len(numbers); i += 2 {
		seeds = append(seeds, rangeSeed{
			start: numbers[i],
			end:   numbers[i] + numbers[i+1] - 1,
		})
	}

	sets := sets(lines)
	for i := 0; i < int(math.Inf(1)); i++ {
		seedSolution := findSeed(i, sets)
		for _, seed := range seeds {
			if seedSolution >= seed.start && seedSolution <= seed.end {
				fmt.Println("Solution B: ", i)
				return
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
