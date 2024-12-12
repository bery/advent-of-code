package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
 * 2024/d2/main.go
 * https://adventofcode.com/2024/day/2
 * need to parse input like 1 2 3 4 to slice of ints
 */
func parseInput(filepath string) [][]int {
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Create a 2D slice to store the numbers
	var numbers [][]int

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read the line
		line := scanner.Text()

		// Split the line into fields (assumes space-separated values)
		fields := strings.Fields(line)

		// Convert each field to an integer
		var row []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return nil
			}
			row = append(row, num)
		}

		// Append the row to the 2D slice
		numbers = append(numbers, row)
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return numbers
}

func evaluateRow(row []int) bool {
	ret := false
	increasing := row[0] < row[1]
	for i := 0; i < len(row)-1; i++ {
		j := i + 1
		if row[i] != row[j] && (row[i] < row[j] && row[j]-row[i] < 4 && increasing || !increasing && row[i] > row[j] && row[i]-row[j] < 4) {
			ret = true
		} else {
			return false
		}
	}
	return ret
}

func main() {
	rows := parseInput("input.txt")
	count := 0
	for _, row := range rows {
		if evaluateRow(row) {
			fmt.Println(row)
			count++
		}
	}
	fmt.Println(count)
}
