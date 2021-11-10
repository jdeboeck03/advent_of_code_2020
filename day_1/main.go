package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// loads input of file into a slice of integer variables
func loadInput(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_int, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		input = append(input, line_int)
	}
	return input, nil
}

func main() {
	filename := "input.txt"
	input, err := loadInput(filename)
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	// Double (Now Triple) for loop that sums every input integer with all integers in input slice except its own
	// Clean Code
	for i, first := range input {
		for j, second := range input {
			for k, third := range input {
				if i != j && i != k && j != k {
					sum := first + second + third
					if sum == 2020 {
						fmt.Printf("First: %d Second: %d Third: %d Sum: %d\n", first, second, third, sum)
						fmt.Printf("VICTORY\n")
						product := first * second * third
						fmt.Printf("Product: %d\n", product)
					}
				}
			}
		}
	}
	sum := input[0] + input[1]
	fmt.Println(sum)
}
