package main

import (
	"log"

	"github.com/GeorgijGrigoriev/minerstattg/pkg/bot"
	c "github.com/GeorgijGrigoriev/minerstattg/pkg/config"
)

func main() {
	log.Println("Starting minerstat")

	conf := c.Config{}

	conf.ReadConfig()
	log.Printf("%v", conf)

	bot.Run("5846901694:AAGm1JwF4qWA-OitiRXxXe6lz405jn5d_VU")
}
