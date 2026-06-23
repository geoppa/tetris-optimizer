package fileutils

import (
	"errors"
	"fmt"
	"os"
)

// CreateSampleFile creates sample.txt with tetrominoes if it does not exist.
// Returns true if the file was newly created, false if it already existed.
func CreateSampleFile() (bool, error) {
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

	// Open the file only if it DOES NOT exist (os.O_EXCL)
	file, err := os.OpenFile("sample.txt", os.O_CREATE|os.O_WRONLY|os.O_EXCL, 0666)
	if err != nil {
		// If the file already exists, return false with no error
		if errors.Is(err, os.ErrExist) {
			return false, nil
		}
		return false, fmt.Errorf("fail to create file: %w", err)
	}
	defer file.Close()

	// Write the tetromino content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return false, fmt.Errorf("fail writing to file: %w", err)
	}

	return true, nil
}
