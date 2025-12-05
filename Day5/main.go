package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Min, Max int
}

func main() {

	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	ranges := []Range{}
	partOneCount := 0
	partTwoCount := 0

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-") {
			rangeParts := strings.Split(line, "-")
			min, _ := strconv.Atoi(rangeParts[0])
			max, _ := strconv.Atoi(rangeParts[1])
			ranges = append(ranges, Range{min, max})
			continue
		}
		if line != "" {
			n, _ := strconv.Atoi(line)
			if isInAnyRange(n, ranges) {
				partOneCount++
			}
		}

	}

	partTwoCount = totalNonOverlappingRanges(ranges)

	fmt.Printf("Part 1 Sum: %d\n", partOneCount)
	fmt.Printf("Part 2 Sum: %d\n", partTwoCount)

}

func isInAnyRange(n int, ranges []Range) bool {
	for _, r := range ranges {
		if n >= r.Min && n <= r.Max {
			return true
		}
	}
	return false
}

func totalNonOverlappingRanges(ranges []Range) int {
	if len(ranges) == 0 {
		return 0
	}

	// Sort ranges by Min value
	sorted := make([]Range, len(ranges))
	copy(sorted, ranges)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Min < sorted[j].Min
	})

	// Merge overlapping ranges
	merged := []Range{sorted[0]}
	for i := 1; i < len(sorted); i++ {
		last := &merged[len(merged)-1]
		current := sorted[i]

		// If current overlaps or is adjacent to last, merge them
		if current.Min <= last.Max+1 {
			if current.Max > last.Max {
				last.Max = current.Max
			}
		} else {
			// No overlap, add as new range
			merged = append(merged, current)
		}
	}

	// Sum up the sizes of merged ranges
	total := 0
	for _, r := range merged {
		total += r.Max - r.Min + 1
	}

	return total
}
