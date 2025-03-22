package main

import (
	"calendar/painter"
	"calendar/solver"
	"calendar/tgbot"
	"fmt"
	_ "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

func main() {

	field := solver.InitField()

	var month, day string

	mode := os.Getenv("RUN_MODE")

	if mode == "tg" {
		tgbot.RunBot(field, month, day)
	}

	if mode == "local" {
		figures := solver.InitFigures()

		fmt.Printf("Input mounth: ")
		fmt.Scanf("%s", &month)
		fmt.Printf("Input day: ")
		fmt.Scanf("%s", &day)

		newField, _ := solver.SolvePuzzle(field, figures, month, "["+day+"]")

		painter.Draw(newField)

		fmt.Println("Solved")

	}

}
