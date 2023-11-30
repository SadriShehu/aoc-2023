package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	calories := 0
	maxCalories := 0
	elfID := 0

	for _, line := range lines {
		snack, err := strconv.Atoi(line)
		if err != nil {
			elfID++
			if calories > maxCalories {
				maxCalories = calories
			}
			// Start with a new elf
			calories = 0
			continue
		}
		calories += snack
	}

	fmt.Println("Elf", elfID, "has max calories", maxCalories)
}
