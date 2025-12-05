package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	dialPosition := 50
	zeroCount := 0

	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()

		// Check if line is not empty
		if len(line) == 0 {
			continue
		}

		// Get the first character
		firstChar := line[0]

		// Extract the number part (everything after the first character)
		numberStr := line[1:]

		// Parse the number as an integer
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			fmt.Printf("Error parsing number from line '%s': %v\n", line, err)
			continue
		}
		if firstChar == 'L' {
			number *= -1
		}
		initialPosition := dialPosition
		dialPosition += number
		zeroCount += int(math.Floor(math.Abs(float64(dialPosition) / 100)))

		if initialPosition*dialPosition < 0 {
			zeroCount++
		}
		if dialPosition == 0 {
			zeroCount++
		}
		dialPosition = ((dialPosition % 100) + 100) % 100
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println("Zero count:", zeroCount)
}
