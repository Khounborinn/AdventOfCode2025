package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	line := strings.TrimSpace(string(content))
	parts := strings.Split(line, ",")

	sum := 0

	for _, part := range parts {
		rangeParts := strings.Split(part, "-")

		leftStr := rangeParts[0]
		rightStr := rangeParts[1]

		left, _ := strconv.Atoi(leftStr)

		right, _ := strconv.Atoi(rightStr)

		for i := left; i <= right; i++ {

			numDigits := int(math.Floor(math.Log10(float64(i)))) + 1
			if numDigits%2 != 0 {
				continue
			}

			firstHalf, secondHalf := splitNumber(i)
			if firstHalf == secondHalf {
				sum += i
			}
		}

	}
	fmt.Printf("\nPart 1 Sum: %d\n", sum)
}

func part2() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	line := strings.TrimSpace(string(content))
	parts := strings.Split(line, ",")

	sum := 0

	for _, part := range parts {
		rangeParts := strings.Split(part, "-")

		leftStr := rangeParts[0]
		rightStr := rangeParts[1]

		left, _ := strconv.Atoi(leftStr)

		right, _ := strconv.Atoi(rightStr)

		for i := left; i <= right; i++ {
			numDigits := int(math.Floor(math.Log10(float64(i)))) + 1

			sum += enhancedSplitNumber(i, numDigits)
		}

	}
	fmt.Printf("\nPart 2 Sum: %d\n", sum)
}

func splitNumber(n int) (int, int) {
	if n == 0 {
		return 0, 0
	}

	numDigits := int(math.Floor(math.Log10(float64(n)))) + 1

	splitPoint := numDigits / 2

	divisor := int(math.Pow(10, float64(numDigits-splitPoint)))

	firstHalf := n / divisor
	secondHalf := n % int(math.Pow(10, float64(splitPoint)))

	return firstHalf, secondHalf
}

func enhancedSplitNumber(n int, numDigits int) int {
	if n == 0 {
		return 0
	}

	for partitionSize := 1; partitionSize <= numDigits/2; partitionSize++ {
		if numDigits%partitionSize != 0 {
			continue
		}

		numPartitions := numDigits / partitionSize
		if numPartitions < 2 {
			continue
		}

		firstPartition := n / int(math.Pow(10, float64(numDigits-partitionSize)))

		allMatch := true
		temp := n

		for i := 0; i < numPartitions; i++ {
			currentPartition := temp % int(math.Pow(10, float64(partitionSize)))

			if currentPartition != firstPartition {
				allMatch = false
				break
			}
			temp = temp / int(math.Pow(10, float64(partitionSize)))
		}

		if allMatch {
			return n
		}
	}

	return 0
}
