package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

// loads input of file into a slice of integer variables
func loadInput() ([]string, error) {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		input = append(input, line)
	}
	return input, nil
}

// process fields, so extract all fields of every passport and put it in slice
func processFields(input []string) []string {
	var fields = make([]string, len(input))
	return fields
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	fmt.Println(input)
	fields := processFields(input)
	fmt.Println(fields[0])
	fmt.Printf("Time: %s\n", time.Since(start1))
}
