package main

import (
	"fmt"
	"os"
	"strings"
	"tetris-optimizer/fileutils"
	"tetris-optimizer/solver"
)

func main() {
	// Check command-line arguments. Must be exactly 2.
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [filename]")
		return
	}

	fileName := os.Args[1]

	// Call the external function to create sample.txt if requested.
	if fileName == "sample.txt" {
		created, err := fileutils.CreateSampleFile()
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			return
		}

		// Print the appropriate message based on whether the file was created or already existed
		if created {
			fmt.Println("The file sample.txt with the tetrominoes was successfully created.")
		} else {
			fmt.Println("The file sample.txt already exists. Skipping creation.")
		}
	}

	// Read the raw content of the input file.
	rawContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("ERROR: File not found")
		return
	}

	// Parse and validate the tetrominoes from the file content.
	rawBlocks, err := parseAndValidate(string(rawContent))
	if err != nil {
		fmt.Println("ERROR: Bad file content")
		return
	}

	// Convert raw string blocks into structured Tetrominoes.
	tetrominoes := fileutils.ConvertToTetrominoes(rawBlocks)

	// Calculate the initial minimum square size needed.
	size := solver.GetMinimumSize(len(tetrominoes))

	// Infinite loop that increases the board size until a solution is found
	for {
		board := solver.CreateBoard(size)

		// Try to solve the board starting with the first tetromino (index 0)
		if board.Solve(tetrominoes, 0) {
			// Print the final solved board and exit the program
			board.Print()
			return
		}

		// If no solution is found for the current size, increment the size and try again
		size++
	}
}

// parseAndValidate splits the file into individual tetrominoes and validates their structure.
func parseAndValidate(content string) ([]string, error) {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	blocks := strings.Split(content, "\n\n")
	var validTetrominoes []string

	for i, block := range blocks {
		block = strings.TrimSpace(block)
		if block == "" {
			if i == len(blocks)-1 {
				continue
			}
			return nil, fmt.Errorf("invalid empty block")
		}

		lines := strings.Split(block, "\n")
		if len(lines) != 4 {
			return nil, fmt.Errorf("each tetromino must be exactly 4 lines high")
		}

		sharpCount := 0
		for _, line := range lines {
			if len(line) != 4 {
				return nil, fmt.Errorf("each line must be exactly 4 characters wide")
			}
			for _, ch := range line {
				if ch == '#' {
					sharpCount++
				} else if ch != '.' {
					return nil, fmt.Errorf("invalid character found: %c", ch)
				}
			}
		}

		if sharpCount != 4 {
			return nil, fmt.Errorf("tetromino must contain exactly 4 '#' characters")
		}

		if !isValidShape(lines) {
			return nil, fmt.Errorf("invalid tetromino shape")
		}

		validTetrominoes = append(validTetrominoes, block)
	}

	if len(validTetrominoes) == 0 || len(validTetrominoes) > 26 {
		return nil, fmt.Errorf("invalid number of tetrominoes")
	}

	return validTetrominoes, nil
}

// isValidShape counts the total side connections between '#' blocks.
func isValidShape(lines []string) bool {
	connections := 0
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if lines[r][c] == '#' {
				if r > 0 && lines[r-1][c] == '#' {
					connections++
				}
				if r < 3 && lines[r+1][c] == '#' {
					connections++
				}
				if c > 0 && lines[r][c-1] == '#' {
					connections++
				}
				if c < 3 && lines[r][c+1] == '#' {
					connections++
				}
			}
		}
	}
	return connections == 6 || connections == 8
}
