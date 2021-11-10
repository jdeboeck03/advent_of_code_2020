package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// loads input of file into a slice of integer variables
func loadInput() ([]string, error) {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}
	return input, nil
}

func createArea(input []string) ([][]string, error) {
	// Count width and height of area
	var height = len(input)
	var width = len(input[0])
	var area = make([][]string, height)
	for i := range area {
		area[i] = make([]string, width)
	}

	// Loop through input
	// Put each character as an element in area slice
	var index = 0
	for _, line := range input {
		chars := []rune(line)
		for j, char := range chars {
			area[index][j] = string(char)
		}
		// fmt.Println(len(chars))
		// fmt.Println(area[index][0])
		index += 1
	}
	return area, nil
}

func findTrees(area [][]string) (int, error) {
	var tree_cnt = 0
	var tree = "#"
	var width = len(area[0])
	var height = len(area)
	fmt.Println("Size of Area:")
	fmt.Printf("Height: %d\n", height)
	fmt.Printf("Width: %d\n", width)
	var height_index = 0
	var width_index = 0
	var right_cnt = 3
	if area[height_index][width_index] == tree {
		tree_cnt += 1
	}
	// As long as we have not reached the bottom of the area
	// continue looping
	for height_index < height-1 {
		// Descend downwards
		height_index += 1
		width_index += right_cnt
		right_cnt = 3
		// Check if there is a tree
		// fmt.Printf("Height: %d Width: %d\n", height_index, width_index)
		// fmt.Println(area[height_index][width_index])
		if area[height_index][width_index] == tree {
			// Found a tree
			//fmt.Println("BOINK")
			//fmt.Printf("Height: %d Width: %d\n", height_index, width_index)
			//fmt.Println(area[height_index][width_index])
			tree_cnt += 1
		}

		// Check if we reached the border of the area
		// If so we need to start from the left
		var border_check = width_index + right_cnt
		if border_check >= width {
			// fmt.Println("Reached the end")
			// Check how far we still need to go
			right_cnt = border_check - width
			width_index = 0
		}
	}
	return tree_cnt, nil
}

func main() {
	input, err := loadInput()
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	area, err := createArea(input)
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	result, err := findTrees(area)
	if err != nil {
		log.Fatalf("Program failed: %s", err)
	}
	fmt.Printf("Tree count: %d", result)
}
