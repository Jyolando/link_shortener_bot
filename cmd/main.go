package main

import (
	"log"

	"github.com/Jyolando/link_shortener_bot/pkg/telegram"
	"github.com/Jyolando/link_shortener_bot/pkg/web"
)

func main() {
	log.Println("Starting link shortener bot...")
	log.Println("Initializing telegram bot...")
	go telegram.StartBot("5417497178:AAE5G2KhNMWCOBYlRiNTvk6NxWCYire2vws")
	web.Init()
}
