package commands

import (
	"database/sql"
	"log"
	"math/rand"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type link string

func (str link) IsUrl() bool {
	u, err := url.Parse(string(str))
	return err == nil && u.Scheme != "" && u.Host != ""
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CmdShorten(bot *tgbotapi.BotAPI, update tgbotapi.Update) int {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.Text = "Please, enter your link. (Example: https://google.ru)"
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Println(err)
		msg.Text = "Error while sending message"
		return 0
	}
	return 3 // attempt to shorten link
}

func TryShort(bot *tgbotapi.BotAPI, update tgbotapi.Update, attempt int, db *sql.DB) bool {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if link(update.Message.Text).IsUrl() {
		rand := RandStringRunes(6)

		log.Println("Shortening link:", update.Message.Text)
		msg.Text = "Link " + update.Message.Text + " is shortened to jyolando.ru/" + rand
		bot.Send(msg)
		_, err := db.Query("INSERT INTO links (fulllink, shortlink) VALUES ($1, $2)", update.Message.Text, rand)
		if err != nil {
			panic(err)
		}
		CmdStart(bot, update)
		return true
	} else {
		log.Println("Invalid link:", update.Message.Text)
		if attempt != 1 {
			msg.Text = "Please, enter valid link. (Example: https://google.ru)"
			bot.Send(msg)
		} else {
			msg.Text = "You have exceeded the limit of attempts. Please, try again."
			bot.Send(msg)
			CmdStart(bot, update)
		}
		return false

	}
}
