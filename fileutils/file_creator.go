package fileutils

import (
	"fmt"
	"os"
)

// creates a sample.txt with the tetrominoes
func CreateSampleFile() error {
	content := `...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

....
.##.
.##.
....

....
....
##..
.##.

##..
.#..
.#..
....

....
###.
.#..
....
`

	file, err := os.Create("sample.txt")
	if err != nil {
		return fmt.Errorf("Fail to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("Fail writing to file: %w", err)
	}

	return nil
}
