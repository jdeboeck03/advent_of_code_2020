package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

// process fields, so find the valid passports
func processFields(input []string) int {
	field_checker := make(map[string]bool)
	req_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "ecl", "pid", "cid"}
	// init field checker
	for _, req_field := range req_fields {
		field_checker[req_field] = false
	}
	var valid_passports = 0
	var valid_passport = true
	for i, line := range input {
		// Reached a new passport so req_fields need to be resetted
		if line == "" || i == len(input) {
			fmt.Println("New Passport!")
			fmt.Println(field_checker)
			for _, req_field := range req_fields {
				if !field_checker[req_field] && req_field != "cid" {
					// one of the fields on passport is not filled in
					fmt.Println("Oei Oei")
					valid_passport = false
				}
			}
			// reset req fields
			for _, req_field := range req_fields {
				field_checker[req_field] = false
				//fmt.Println(field_checker)
			}
			// Check if passport is valid
			if valid_passport {
				valid_passports += 1
			}
			valid_passport = true
		} else {
			// Now we need to process the entire field (use a regex and substring)
			re := regexp.MustCompile("[: ]")
			line_split := re.Split(line, -1)
			//fmt.Println(line_split)
			for i, field := range line_split {
				//fmt.Println(field)
				// Only check even indexes
				if i%2 == 0 {
					//fmt.Println(field)
					field_checker[field] = true
				}
			}
		}

	}
	// Check final passport
	fmt.Println("New Passport!")
	fmt.Println(field_checker)
	for _, req_field := range req_fields {
		if !field_checker[req_field] && req_field != "cid" {
			// one of the fields on passport is not filled in
			fmt.Println("Oei Oei")
			valid_passport = false
		}
	}
	// reset req fields
	for _, req_field := range req_fields {
		field_checker[req_field] = false
		//fmt.Println(field_checker)
	}
	// Check if passport is valid
	if valid_passport {
		valid_passports += 1
	}
	valid_passport = true
	return valid_passports
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	start1 := time.Now()
	fmt.Println(input[0])
	fields := processFields(input)
	fmt.Println(fields)
	fmt.Printf("Time: %s\n", time.Since(start1))
}
