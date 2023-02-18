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

	bot.Run(conf.Token)
}
