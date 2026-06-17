package solver

import (
	"fmt"
	"math"
)

// Point represents a single block coordinate (X, Y).
type Point struct {
	X, Y int
}

// Tetromino represents a validated and normalized Tetris piece.
type Tetromino struct {
	ID     rune
	Blocks [4]Point // Διορθώθηκε: Ορίστηκε σωστά ως πίνακας 4 στοιχείων τύπου Point
}

// Board represents the grid where tetrominoes will be placed.
type Board [][]rune

// GetMinimumSize calculates the starting square size based on the number of pieces.
func GetMinimumSize(numPieces int) int {
	area := numPieces * 4
	side := math.Ceil(math.Sqrt(float64(area)))
	return int(side)
}

// CreateBoard initializes a square board of a given size filled with '.' characters.
func CreateBoard(size int) Board {
	board := make(Board, size)
	for i := range board {
		board[i] = make([]rune, size)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	return board
}

// Print displays the current state of the board.
func (b Board) Print() {
	for i := range b {
		for j := range b[i] {
			fmt.Printf("%c", b[i][j])
		}
		fmt.Println()
	}
}

// CanPlace checks if a tetromino can be placed on the board at a specific (row, col) anchor.
func (b Board) CanPlace(t Tetromino, row, col int) bool {
	boardSize := len(b)

	for _, block := range t.Blocks {
		// Calculate target coordinates on the board
		targetRow := row + block.Y
		targetCol := col + block.X

		// Check if the block goes out of the board boundaries
		if targetRow >= boardSize || targetCol >= boardSize || targetRow < 0 || targetCol < 0 {
			return false
		}

		// Check if the position is already occupied by another tetromino
		if b[targetRow][targetCol] != '.' {
			return false
		}
	}

	return true
}

// Place stamps the tetromino's ID onto the board at the given (row, col) anchor.
func (b Board) Place(t Tetromino, row, col int) {
	for _, block := range t.Blocks {
		b[row+block.Y][col+block.X] = t.ID
	}
}

// Remove clears the tetromino from the board by replacing its ID with '.' characters.
func (b Board) Remove(t Tetromino, row, col int) {
	for _, block := range t.Blocks {
		b[row+block.Y][col+block.X] = '.'
	}
}

// Solve uses recursion and backtracking to find the correct placement for all tetrominoes.
func (b Board) Solve(tetrominoes []Tetromino, index int) bool {
	// Base Case: If all tetrominoes are successfully placed, we are done!
	if index == len(tetrominoes) {
		return true
	}

	boardSize := len(b)
	currentPiece := tetrominoes[index]

	// Iterate through every cell of the board to find a valid spot
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {

			// Check if the current tetromino can fit at the current (row, col) anchor
			if b.CanPlace(currentPiece, row, col) {

				// Action: Place the piece on the board
				b.Place(currentPiece, row, col)

				// Recursion: Try to place the next tetromino (index + 1)
				if b.Solve(tetrominoes, index+1) {
					return true // Solution found down this path!
				}

				// Backtrack: If the next pieces didn't fit, remove this piece and try the next cell
				b.Remove(currentPiece, row, col)
			}
		}
	}

	// If the piece cannot be placed anywhere on the current board, return false
	return false
}
