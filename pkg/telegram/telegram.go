package telegram

import (
	"log"

	"github.com/Jyolando/link_shortener_bot/pkg/database"
	"github.com/Jyolando/link_shortener_bot/pkg/telegram/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var waitLink int

func StartBot(key string) {
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Telegram Bot online, nickname:", bot.Self.UserName)
	log.Println("Initializing database...")
	database.Init()
	getUpdates(bot)
}

func getUpdates(bot *tgbotapi.BotAPI) {
	updates := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))

	for update := range updates {
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		if update.Message == nil {
			continue
		}

		if update.Message.Command() == "start" {
			commands.CmdStart(bot, update)
		}

		if waitLink != 0 {
			if commands.TryShort(bot, update, waitLink) {
				waitLink = 0 // if user entered valid link, reset waitLink
				log.Println("waitLink reset")
				continue
			}
			waitLink-- // if user send invalid link, waitLink will be decremented
		}

		switch update.Message.Text {
		case "Shorten link":
			waitLink = commands.CmdShorten(bot, update)
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
