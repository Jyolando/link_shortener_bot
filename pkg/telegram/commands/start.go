package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var startKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Shorten link"),
	), tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Your links"),
	),
)

func CmdStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	if update.Message.Chat.UserName == "" {
		msg.Text = "Please, set your username in Telegram settings"
	} else {
		msg.Text = "Welcome, " + update.Message.Chat.UserName + "!"
	}
	msg.Text = msg.Text + "\nHello, I'm a link shortener bot.\nUse bottom keyboard to use me."
	msg.ReplyMarkup = startKeyboard
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}
