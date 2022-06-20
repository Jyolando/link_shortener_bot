package main

import (
	"log"

	"github.com/Jyolando/link_shortener_bot/pkg/telegram"
)

func main() {
	log.Println("Starting link shortener bot...")
	log.Println("Initializing telegram bot...")
	telegram.StartBot("")
}
