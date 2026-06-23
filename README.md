# Tetris Optimizer

A high-performance Go program that takes a list of Tetrominoes from a text file and arranges them into the smallest possible square board using a backtracking algorithm.

## Features
* **Optimal Square Finder**: Dynamically calculates and grows the minimum board size required.
* **Strict Validation**: Validates Tetromino shapes, connection constraints, grid sizes, and valid characters.
* **Automatic Sample Generation**: Built-in feature to generate a sample input file if needed.
* **Support for 1-26 Blocks**: Handles up to 26 Tetrominoes, assigning a unique uppercase letter (A-Z) to each.

---

## Installation

### Prerequisites
* **Go** installed on your system (version 1.16 or higher recommended).

### Setup
Clone this repository or navigate to your project root folder where the `main.go` file is located.

---

## Usage

### Running the Program
Execute the program by passing the path to your input file as a command-line argument:

```bash
go run . <filename>
```

### Example
```bash
go run . sample.txt
```

### Generate a Sample File
If you pass `sample.txt` as an argument and it does not exist, the program will automatically generate it for you:
```bash
go run . sample.txt
```
*Output:* `The file sample.txt with the tetrominoes was successfully created.`

---

## Input File Format

The input file must follow strict formatting rules:
1. Each Tetromino must fit inside a **4x4 text grid**.
2. It must consist *only* of `.` (empty space) and `#` (block piece).
3. Exactly **4 `#` characters** must exist per block, forming a valid connected Tetris shape.
4. Each block must be separated by **exactly one empty line** (`\n\n`).

### Input Example (`sample.txt`)
```text
#...
#...
#...
#...

..#.
..#.
..#.
..#.

####
....
....
....
```

---

## Output Format

The program identifies each block sequentially using uppercase letters (`A` for the first block, `B` for the second, etc.). It outputs the smallest possible square grid filled with these blocks.

### Output Example
```text
ABBBA
AB..A
AB..A
AB..A
A....
```

---

## Error Handling

The program handles bad inputs gracefully and prints specific error messages:
* `Usage: go run . [filename]` — Triggered if the filename argument is missing or extra arguments are given.
* `ERROR: File not found` — Triggered if the specified file does not exist.
* `ERROR: Bad file content` — Triggered if formatting, block connections, character types, or block counts (must be between 1 and 26) are invalid.
