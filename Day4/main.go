package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a 2D matrix (slice of byte slices)
	var matrix [][]byte

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Convert line to byte slice and append to matrix
		matrix = append(matrix, []byte(line))
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	partOne(matrix)
	partTwo(matrix)
}

func partOne(matrix [][]byte) {
	sum := 0
	for i := range matrix {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '@' {
				neighbors := countAdjacentAt(matrix, i, j)
				if neighbors < 4 {
					sum++
				}
			}
		}
	}
	fmt.Printf("Part 1 Sum: %d\n", sum)
}

func partTwo(matrix [][]byte) {
	sum := 0
	iterate := true
	for iterate {
		iterationSum := 0
		for i := range matrix {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == '@' {
					neighbors := countAdjacentAt(matrix, i, j)
					if neighbors < 4 {
						matrix[i][j] = '.'
						iterationSum++
					}
				}
			}
		}
		if iterationSum == 0 {
			iterate = false
		}
		sum += iterationSum
	}

	fmt.Printf("Part 2 Sum: %d\n", sum)
}

func countAdjacentAt(matrix [][]byte, row, col int) int {
	count := 0
	rows := len(matrix)
	cols := len(matrix[0])

	directions := [][]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up, left
		{-1, 1},  // up, right
		{1, -1},  // down, left
		{1, 1},   // down, right
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if matrix[newRow][newCol] == '@' {
				count++
			}
		}
	}

	return count
}
