package bot

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/GeorgijGrigoriev/minerstattg/pkg/db"
	a "github.com/GeorgijGrigoriev/minerstattg/pkg/http"
	"github.com/GeorgijGrigoriev/minerstattg/pkg/templates"

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
		var d = db.Init()

		d.Get()

		isAdded, err := d.CheckUser(int(c.Sender().ID))
		if err != nil {
			log.Fatal(err)
		}

		if !isAdded {
			d.AddUser(c.Sender())

			return c.Send(fmt.Sprintf("Welcome %s!", c.Sender().FirstName))
		}
		d.UpdateUser(c.Sender())

		return c.Send(fmt.Sprintf("Welcome back, %s!", c.Sender().FirstName))
	})

	b.Handle("/prices", func(c tele.Context) error {
		pair := c.Message().Payload

		t := a.TradeOgreAPI{
			Pair: "USDT-BTC",
		}

		b := a.BybitAPI{
			Pair: "BTCUSDT",
		}

		wg := sync.WaitGroup{}
		wg.Add(2)

		go b.GetMarket(&wg)
		go t.GetMarket(&wg)

		wg.Wait()

		c.Send(fmt.Sprintf("Prices for %s", pair))
		c.Send(templates.MakeTemplate("TradeOgre", t.Response))
		c.Send(templates.MakeTemplate("ByBit", b.Response))

		return nil
	})

	b.Start()
}
