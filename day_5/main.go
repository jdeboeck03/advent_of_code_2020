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
	input := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		input = append(input, line)
	}
	return input, nil
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	fmt.Println(input)
	start1 := time.Now()
	valid_passports_1 := 5
	fmt.Println("Solution 1:")
	fmt.Printf("Number of valid passports: %d\n", valid_passports_1)
	fmt.Printf("Time: %s\n", time.Since(start1))
}
