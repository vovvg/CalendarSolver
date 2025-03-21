package main

import (
	"calendar/painter"
	"calendar/solver"
	"fmt"
	"log"
)

func main() {

	field := solver.InitField()

	figures := solver.InitFigures()

	var month, day string

	fmt.Printf("Input mounth: ")
	fmt.Scanf("%s", &month)
	fmt.Printf("Input day: ")
	fmt.Scanf("%s", &day)

	newField, _ := solver.SolvePuzzle(field, figures, 0, month, "["+day+"]")

	painter.Draw(newField)

	log.Println("Solved")

}
