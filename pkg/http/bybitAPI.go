package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/GeorgijGrigoriev/minerstattg/internal/api"
)

type BybitAPI struct {
	Response api.Market
	Pair     string
}

func (b *BybitAPI) GetMarket(wg *sync.WaitGroup) error {
	res, err := http.Get(fmt.Sprintf("https://api.bybit.com/v5/market/tickers?category=spot&symbol=%s", b.Pair))
	if err != nil {
		log.Printf("HTTP Call error: %s", err.Error())

		return err
	}
	defer res.Body.Close()

	mrkt := BybitAPIMarkets{}

	json.NewDecoder(res.Body).Decode(&mrkt)

	b.Response.Price = mrkt.Result.List[0].LastPrice
	b.Response.High = mrkt.Result.List[0].High
	b.Response.Low = mrkt.Result.List[0].Low
	b.Response.Volume = mrkt.Result.List[0].Volume

	wg.Done()

	return nil
}
