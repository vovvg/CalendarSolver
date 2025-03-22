package tgbot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var months = []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

var days = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
	"11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21",
	"22", "23", "24", "25", "26", "27", "28", "29", "30", "31"}

var monthKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("jan"),
		tgbotapi.NewKeyboardButton("feb"),
		tgbotapi.NewKeyboardButton("mar"),
		tgbotapi.NewKeyboardButton("apr"),
		tgbotapi.NewKeyboardButton("may"),
		tgbotapi.NewKeyboardButton("jun"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("jul"),
		tgbotapi.NewKeyboardButton("aug"),
		tgbotapi.NewKeyboardButton("sep"),
		tgbotapi.NewKeyboardButton("oct"),
		tgbotapi.NewKeyboardButton("nov"),
		tgbotapi.NewKeyboardButton("dec"),
	),
)

var dayKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
		tgbotapi.NewKeyboardButton("7"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("8"),
		tgbotapi.NewKeyboardButton("9"),
		tgbotapi.NewKeyboardButton("10"),
		tgbotapi.NewKeyboardButton("11"),
		tgbotapi.NewKeyboardButton("12"),
		tgbotapi.NewKeyboardButton("13"),
		tgbotapi.NewKeyboardButton("14"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("15"),
		tgbotapi.NewKeyboardButton("16"),
		tgbotapi.NewKeyboardButton("17"),
		tgbotapi.NewKeyboardButton("18"),
		tgbotapi.NewKeyboardButton("19"),
		tgbotapi.NewKeyboardButton("20"),
		tgbotapi.NewKeyboardButton("21"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("22"),
		tgbotapi.NewKeyboardButton("23"),
		tgbotapi.NewKeyboardButton("24"),
		tgbotapi.NewKeyboardButton("25"),
		tgbotapi.NewKeyboardButton("26"),
		tgbotapi.NewKeyboardButton("27"),
		tgbotapi.NewKeyboardButton("28"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("29"),
		tgbotapi.NewKeyboardButton("30"),
		tgbotapi.NewKeyboardButton("31"),
	),
)
