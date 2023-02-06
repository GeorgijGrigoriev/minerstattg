package bot

import (
	"fmt"
	"time"

	tele "gopkg.in/telebot.v3"
)

func Run(token string) {
	if token == "" {
		panic("Empty token not allowed")
	}
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 5 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	b.Handle("/start", func(c tele.Context) error {
		fmt.Println(c.Sender().ID)
		fmt.Println(c.Sender().IsPremium)
		fmt.Println(c.Sender().LanguageCode)
		fmt.Println(c.Sender().IsBot)

		return c.Send("Unauthorized")
	})

	b.Start()
}
