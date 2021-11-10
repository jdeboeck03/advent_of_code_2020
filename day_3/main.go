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

func findTrees(area [][]string, right_cnt int, bottom_cnt int) int {
	var tree_cnt = 0
	var tree = "#"
	var width = len(area[0])
	var height = len(area)
	//fmt.Println("Size of Area:")
	//fmt.Printf("Height: %d\n", height)
	//fmt.Printf("Width: %d\n", width)
	var height_index = 0
	var width_index = 0
	var curr_bottom_cnt = bottom_cnt
	var curr_right_cnt = right_cnt
	if area[height_index][width_index] == tree {
		tree_cnt += 1
	}
	// As long as we have not reached the bottom of the area
	// continue looping
	for height_index < height-bottom_cnt {
		// Descend downwards
		height_index += curr_bottom_cnt
		width_index += curr_right_cnt
		curr_right_cnt = right_cnt
		// Check if there is a tree
		//fmt.Printf("Height: %d Width: %d\n", height_index, width_index)
		//fmt.Println(area[height_index][width_index])
		if area[height_index][width_index] == tree {
			// Found a tree
			//fmt.Println("BOINK")
			//fmt.Printf("Height: %d Width: %d\n", height_index, width_index)
			//fmt.Println(area[height_index][width_index])
			tree_cnt += 1
		}

		// Check if we reached the border of the area
		// If so we need to start from the left
		var border_check = width_index + curr_right_cnt
		if border_check >= width {
			//fmt.Println("Reached the end")
			// Check how far we still need to go
			curr_right_cnt = border_check - width
			width_index = 0
		}
	}
	return tree_cnt
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
	start1 := time.Now()
	result1 := findTrees(area, 3, 1)
	fmt.Println("Solution 1: ")
	fmt.Printf("Tree count: %d\n", result1)
	fmt.Printf("Time: %s\n", time.Since(start1))
	start2 := time.Now()
	result2 := findTrees(area, 1, 1)
	result3 := findTrees(area, 5, 1)
	result4 := findTrees(area, 7, 1)
	result5 := findTrees(area, 1, 2)
	total_result := result1 * result2 * result3 * result4 * result5
	fmt.Println("Solution 2: ")
	fmt.Printf("Multiplication Result: %d\n", total_result)
	fmt.Printf("Time: %s\n", time.Since(start2))
}
