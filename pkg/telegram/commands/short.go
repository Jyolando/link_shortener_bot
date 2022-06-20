package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CmdShorten(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Please, enter your link. (Example: https://google.ru)"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
	}
}
