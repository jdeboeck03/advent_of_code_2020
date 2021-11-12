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

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	row_ids := calculateRows(input)
	fmt.Println("Solution 1:")
	fmt.Println(row_ids)
	fmt.Printf("Time: %s\n", time.Since(start1))
}
