package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// loads input of file into a slice of integer variables
func loadInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input, nil
}

// This will need some more split feeling
func main() {
	filename := "input.txt"
	input, err := loadInput(filename)
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	//
	re := regexp.MustCompile("[-, ,:]")
	for _, line := range input {
		input_split := re.Split(line, -1)
		fmt.Println(input_split[4])
	}
}
