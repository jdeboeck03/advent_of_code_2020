package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	var correct_pw_cnt = 0
	re := regexp.MustCompile("[-, ,:]")
	for _, line := range input {
		input_split := re.Split(line, -1)
		lowest_number := input_split[0]
		highest_number := input_split[1]
		character := input_split[2]
		password := input_split[4]
		lowest_number_int, err := strconv.Atoi(lowest_number)
		if err != nil {
			log.Fatalf("Program failed: %s", err)
		}
		highest_number_int, err := strconv.Atoi(highest_number)
		if err != nil {
			log.Fatalf("Program failed: %s", err)
		}

		// Will need to substract one value of each number_int variable, so we can use it for index searching
		// Uses switch case to see if one of the two characters (extracted from index) in de string equals the requested character
		// If both characters match, the password is not valid
		first_index_int := lowest_number_int - 1
		second_index_int := highest_number_int - 1
		switch {
		case string(password[first_index_int]) == character && string(password[second_index_int]) != character:
			correct_pw_cnt += 1
		case string(password[first_index_int]) != character && string(password[second_index_int]) == character:
			//fmt.Printf("Password: %s\n", password)
			//fmt.Printf("Character: %s\n", character)
			//fmt.Printf("First Index: %s Second Index: %s\n", lowest_number, highest_number)
			correct_pw_cnt += 1
		}

		/*
			// Checks if the counted characters lies between the two threshold values
			char_cnt := strings.Count(password, character)
			fmt.Println(lowest_number_int)
			fmt.Println(highest_number_int)
			fmt.Printf("number of chars: %d\n", char_cnt)
			if lowest_number_int <= char_cnt && char_cnt <= highest_number_int {
				fmt.Println("Correct Password Policy")
				correct_pw_cnt += 1
			} else {
				fmt.Println("Wrong Password Policy")
			}*/

	}
	fmt.Printf("Total of valid password policies: %d\n", correct_pw_cnt)
}
