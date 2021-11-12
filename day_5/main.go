package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// loads input of file into a slice of integer variables
func loadInput() ([]string, error) {
	input := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		input = append(input, line)
	}
	return input, nil
}

func calculateRows(input []string) []int {
	row_ids := []int{}
	for _, line := range input {
		fmt.Println(line)
		//Now you can search each boarding pass
		line_chars := strings.Split(line, "")
		var lower_bound = 0
		var upper_bound = 127
		for i := 0; i < 7; i++ {
			char := line_chars[i]
			if char == "F" {
				upper_bound = (lower_bound + upper_bound) / 2
				//fmt.Println(upper_bound)
			} else {
				// char is B
				lower_bound = (lower_bound+upper_bound)/2 + 1
				//fmt.Println(lower_bound)
			}
		}
		row_ids = append(row_ids, lower_bound)
	}
	return row_ids
}

func calculateColumns(input []string) []int {
	column_ids := []int{}
	for _, line := range input {
		fmt.Println(line)
		//Now you can search each boarding pass
		line_chars := strings.Split(line, "")
		var lower_bound = 0
		var upper_bound = 7
		for i := 7; i < 10; i++ {
			char := line_chars[i]
			if char == "L" {
				upper_bound = (lower_bound + upper_bound) / 2
				//fmt.Println(upper_bound)
			} else {
				// char is R
				lower_bound = (lower_bound+upper_bound)/2 + 1
				//fmt.Println(lower_bound)
			}
		}
		column_ids = append(column_ids, lower_bound)
	}
	return column_ids
}

func calculateResults(row_ids []int, column_ids []int) []int {
	results := make([]int, len(row_ids))
	for i, row_id := range row_ids {
		var column_id = column_ids[i]
		result := row_id*8 + column_id
		results[i] = result
	}
	return results
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	row_ids := calculateRows(input)
	column_ids := calculateColumns(input)
	fmt.Println("Solution 1:")
	fmt.Println(row_ids)
	fmt.Println(column_ids)
	seat_ids_1 := calculateResults(row_ids, column_ids)
	fmt.Printf("Result of multiplication: %d\n", seat_ids_1)
	fmt.Printf("Time: %s\n", time.Since(start1))
}
