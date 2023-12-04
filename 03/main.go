package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Number struct {
	Value int
	Start int
	End   int
	Row   int
}

type Point struct {
	X     int
	Y     int
	Value rune
}

var points [][]int = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

const GEAR rune = '*'

func main() {
	res := partOne()
	start1 := time.Now()
	fmt.Printf("Execution time: %s\n", time.Since(start1))
	fmt.Println("Part 1:", res)

	start2 := time.Now()
	res = partTwo()
	fmt.Printf("Execution time: %s\n", time.Since(start2))
	fmt.Println("Part 2:", res)
}

func partOne() int {
	numbers, specialChars := getInput()

	sum := 0
	match := make(map[Number]bool)
	for _, char := range specialChars {
		for _, adj := range points {
			// current ajdacent point
			dx, dy := adj[0], adj[1]
			// points to check
			x, y := char.X+dx, char.Y+dy
			for _, num := range numbers {
				contains := match[Number{Row: num.Row, End: num.End}]
				if !contains && x == num.Row && y >= num.Start && y <= num.End {
					match[Number{Row: num.Row, End: num.End}] = true
					sum += num.Value
				}
			}
		}
	}
	return sum
}

func partTwo() int {
	numbers, specialChars := getInput()
	gears := []Point{}
	for _, v := range specialChars {
		if v.Value == GEAR {
			gears = append(gears, v)
		}
	}

	sum := 0
	match := map[Number]bool{}
	for _, gear := range gears {
		gearList := []Number{}
		for _, adj := range points {
			// current ajdacent point
			dx, dy := adj[0], adj[1]
			// points to check
			x, y := gear.X+dx, gear.Y+dy
			for _, num := range numbers {
				contains := match[Number{Row: num.Row, End: num.End}]
				if !contains && x == num.Row && y >= num.Start && y <= num.End {
					match[Number{Row: num.Row, End: num.End}] = true
					gearList = append(gearList, num)
				}
			}
		}

		if len(gearList) == 2 {
			sum += gearList[0].Value * gearList[1].Value
		}
	}

	return sum
}

func getInput() ([]Number, []Point) {
	numbers := []Number{}
	specialChars := []Point{}

	input := readFile()
	for row, line := range input {
		col := 0
		for col < len(line) {
			if line[col] >= '0' && line[col] <= '9' {
				start := col
				val := ""
				for col < len(line) && line[col] >= '0' && line[col] <= '9' {
					val += string(line[col])
					col++
				}
				num, err := strconv.Atoi(val)
				if err != nil {
					fmt.Printf("Failed to convert %s to number\n", val)
				}
				numbers = append(numbers, Number{
					Value: num,
					Start: start,
					End:   col - 1,
					Row:   row,
				})
			} else if line[col] != '.' {
				specialChars = append(specialChars, Point{
					X:     row,
					Y:     col,
					Value: rune(line[col]),
				})
				col++
			} else {
				col++
			}
		}
	}

	return numbers, specialChars
}

func readFile() []string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(input), "\n")
}
