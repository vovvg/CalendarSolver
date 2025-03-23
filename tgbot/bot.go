package tgbot

import (
	"calendar/painter"
	"calendar/solver"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"slices"
)

func RunBot(field [][]string, month string, day string) {
	figures := solver.InitFigures()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		msgText := update.Message.Text

		if msgText == "/start" {
			msg.Text = "Выберите месяц:"
			msg.ReplyMarkup = monthKeyboard
			bot.Send(msg)
			photoFile := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("img/calendar.png"))
			bot.Send(photoFile)
			continue
		}

		if slices.Contains(months, msgText) {
			month = msgText
			msg.Text = "Вы выбрали месяц: " + month + "\nТеперь выберите день:"
			msg.ReplyMarkup = dayKeyboard
			bot.Send(msg)
		} else if slices.Contains(days, msgText) {
			day = msgText

			newField, _ := solver.SolvePuzzle(field, figures, month, "["+day+"]")
			painter.Draw(newField)

			month, day = "", ""

			photoFile := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("calendar_solved.png"))
			msg.ReplyMarkup = monthKeyboard
			bot.Send(photoFile)
			msg.Text = "Пазл решён! Выберите новый месяц:"
			bot.Send(msg)
			photoFile = tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath("img/calendar.png"))
			bot.Send(photoFile)
		} else {
			month, day = "", ""
			msg.Text = "Ошибка. Выберите месяц:"
			msg.ReplyMarkup = monthKeyboard
			bot.Send(msg)
		}
	}
}
