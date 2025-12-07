package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a matrix to store the data
	var matrix [][]int
	var operators []string

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		hasNumber := strings.ContainsAny(line, "0123456789")

		if !hasNumber {
			fields := strings.Fields(line)
			operators = append(operators, fields...)
			continue
		}

		fields := strings.Fields(line)
		row := make([]int, 0, len(fields))

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Printf("Error converting %s to int: %v\n", field, err)
				continue
			}
			row = append(row, num)
		}

		if len(row) > 0 {
			matrix = append(matrix, row)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	if len(matrix) > 0 {
		numCols := len(matrix[0])
		transposed := make([][]int, numCols)

		for i := 0; i < numCols; i++ {
			transposed[i] = make([]int, len(matrix))
			for j := 0; j < len(matrix); j++ {
				transposed[i][j] = matrix[j][i]
			}
		}
		matrix = transposed
	}

	total := 0
	for i, row := range matrix {
		total += solve(row, operators[i])
	}

	fmt.Printf("Total: %d\n", total)

}

func solve(n []int, operand string) int {
	var result int
	if operand == "*" {
		result = 1
	} else {
		result = 0
	}
	for _, i := range n {
		if operand == "*" {
			result *= i
		} else {
			result += i
		}
	}
	return result
}
