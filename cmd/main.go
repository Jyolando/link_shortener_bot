package main

import (
	"fmt"

	"github.com/Jyolando/link_shortener_bot/pkg/database"
)

func main() {
	DB := database.Init()
	fmt.Println(DB.String())
}
