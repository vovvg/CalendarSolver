package main

import (
	"fmt"
)

func main() {

	field := initField()

	figures := initFigures()

	font := LoadFont()

	var month, day string

	fmt.Printf("Input mounth: ")
	fmt.Scanf("%s", &month)
	fmt.Printf("Input day: ")
	fmt.Scanf("%s", &day)

	newField, _ := solvePuzzle(field, figures, 0, month, "["+day+"]")

	Draw(newField, font)

	fmt.Println("Solved puzzle: ")
	for _, row := range newField {
		fmt.Println(row)
	}
}
