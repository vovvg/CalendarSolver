package main

import (
	"fmt"
	"strings"
)

func main() {

	field := [][]string{
		{"jan", "feb", "mar", "apr", "may", "jun", "[x]"},
		{"jul", "aug", "sep", "oct", "nov", "dec", "[x]"},
		{"[1]", "[2]", "[3]", "[4]", "[5]", "[6]", "[7]"},
		{"[8]", "[9]", "[10]", "[11]", "[12]", "13]", "[14]"},
		{"[15]", "[16]", "[17]", "[18]", "[19]", "[20]", "[21]"},
		{"[22]", "[23]", "[24]", "[25]", "[26]", "[27]", "[28]"},
		{"[29]", "[30]", "[31]", "[x]", "[x]", "[x]", "[x]"}}

	figures := initFigures()

	var month, day string

	for _, row := range field {
		fmt.Println(row)
	}

	fmt.Printf("Input mounth: ")
	fmt.Scanf("%s", &month)
	fmt.Printf("Input day: ")
	fmt.Scanf("%s", &day)

	newField, _ := solvePuzzle(field, figures, 0, month, "["+day+"]")

	fmt.Println("Solved puzzle: ")
	for _, row := range newField {
		fmt.Println(row)
	}
}

func solvePuzzle(field [][]string, figures []*figure, idx int, month string, date string) ([][]string, bool) {
	if idx >= len(figures) {
		return field, true
	}

	for row := 0; row < len(field); row++ {
		for col := 0; col < len(field[0]); col++ {
			for pos := 0; pos < len(figures[idx].position); pos++ {
				if canInsert(field, figures[idx].position[pos], row, col, month, date) {
					exField := deepCopy(field)

					insertPattern(field, figures[idx].position[pos], row, col)
					figures[idx].placed = true

					if newField, success := solvePuzzle(field, figures, idx+1, month, date); success {
						return newField, true
					}

					field = deepCopy(exField)
					figures[idx].placed = false
				}
			}
		}
	}

	return field, false
}

func canInsert(field [][]string, pattern [][]string, startRow, startCol int, month string, date string) bool {
	rows, cols := len(field), len(field[0])
	patternRows, patternCols := len(pattern), len(pattern[0])

	if startRow+patternRows > rows || startCol+patternCols > cols {
		return false
	}

	for i := 0; i < patternRows; i++ {
		for j := 0; j < patternCols; j++ {
			cell := field[startRow+i][startCol+j]
			patternValue := pattern[i][j]

			if (cell == month || cell == date || cell == "[x]" || strings.Contains(cell, "{")) && patternValue != "{x}" {
				return false
			}
		}
	}

	return true
}

func insertPattern(field [][]string, pattern [][]string, startRow, startCol int) {
	for i := 0; i < len(pattern); i++ {
		for j := 0; j < len(pattern[i]); j++ {
			if pattern[i][j] != "{x}" {
				field[startRow+i][startCol+j] = pattern[i][j]
			}
		}
	}
}

func deepCopy(field [][]string) [][]string {
	copyField := make([][]string, len(field))
	for i := range field {
		copyField[i] = make([]string, len(field[i]))
		copy(copyField[i], field[i])
	}
	return copyField
}
