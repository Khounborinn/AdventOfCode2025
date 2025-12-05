package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	partOneSum := 0
	partTwoSum := 0
	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		partOneSum += largestCombination(line)
		partTwoSum += enhancedLargestCombination(line)
	}
	fmt.Printf("Part 1 Sum: %d\n", partOneSum)
	fmt.Printf("Part 1 Sum: %d\n", partTwoSum)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
}

func largestCombination(line string) int {
	if len(line) < 2 {
		return 0
	}

	maxCombo := 0
	maxRight := make([]int, len(line))
	maxDigit := -1
	for i := len(line) - 1; i >= 0; i-- {
		maxRight[i] = maxDigit
		digit := int(line[i] - '0')
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	for i := 0; i < len(line)-1; i++ {
		if maxRight[i] >= 0 {
			combo := int(line[i]-'0')*10 + maxRight[i]
			if combo > maxCombo {
				maxCombo = combo
			}
		}
	}

	return maxCombo
}

func enhancedLargestCombination(line string) int {
	if len(line) < 12 {
		return 0
	}
	if len(line) == 12 {
		n, _ := strconv.Atoi(line)
		return n
	}

	return 0
}
