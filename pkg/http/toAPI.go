package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/GeorgijGrigoriev/minerstattg/internal/api"
)

type TradeOgreAPI struct {
	Response api.Market
	Pair     string
}

func (t *TradeOgreAPI) GetMarket(wg *sync.WaitGroup) error {
	resp, err := http.Get(fmt.Sprintf("https://tradeogre.com/api/v1/ticker/%s", t.Pair))
	if err != nil {
		log.Printf("API call error: %s", err.Error())

		return err
	}
	defer resp.Body.Close()

	mrkt := TradeOgreAPIMarkets{}
	json.NewDecoder(resp.Body).Decode(&mrkt)

	t.Response.Price = mrkt.Price
	t.Response.High = mrkt.High
	t.Response.Low = mrkt.Low
	t.Response.Volume = mrkt.Volume

	wg.Done()

	return nil
}
