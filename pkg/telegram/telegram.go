package telegram

import (
	"log"
	"net/url"

	"github.com/Jyolando/link_shortener_bot/pkg/database"
	"github.com/Jyolando/link_shortener_bot/pkg/telegram/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var waitLink bool = false

func StartBot(key string) {
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Telegram Bot online, nickname:", bot.Self.UserName)
	log.Println("Initializing database...")
	DB := database.Init()
	log.Println(DB.String())
	getUpdates(bot)
}

func getUpdates(bot *tgbotapi.BotAPI) {
	updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

	for update := range updates {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message == nil {
			continue
		}

		if update.Message.Command() == "start" {
			commands.CmdStart(bot, update)
		}

		if waitLink {
			if IsUrl(update.Message.Text) {
				log.Println("Shortening link:", update.Message.Text)
				msg.Text = "Link " + update.Message.Text + " is shortened to KEKW"
				waitLink = false
			} else {
				msg.Text = "Please, enter valid link. (Example: https://google.ru)"
			}
			bot.Send(msg)
		}

		switch update.Message.Text {
		case "Shorten link":
			commands.CmdShorten(bot, update)
			waitLink = true
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
