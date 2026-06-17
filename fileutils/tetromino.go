package fileutils

import (
	"strings"
	"tetris-optimizer/solver" // Import the solver package to use its structs
)

// ConvertToTetrominoes transforms raw string blocks into normalized solver.Tetromino structs.
func ConvertToTetrominoes(rawBlocks []string) []solver.Tetromino {
	var list []solver.Tetromino

	for i, block := range rawBlocks {
		var t solver.Tetromino
		t.ID = rune('A' + i)

		lines := strings.Split(block, "\n")
		minX, minY := 4, 4
		blockIndex := 0

		for r := 0; r < 4; r++ {
			for c := 0; c < 4; c++ {
				if lines[r][c] == '#' {
					t.Blocks[blockIndex] = solver.Point{X: c, Y: r}
					blockIndex++

					if c < minX {
						minX = c
					}
					if r < minY {
						minY = r
					}
				}
			}
		}

		for b := 0; b < 4; b++ {
			t.Blocks[b].X -= minX
			t.Blocks[b].Y -= minY
		}

		list = append(list, t)
	}

	return list
}
