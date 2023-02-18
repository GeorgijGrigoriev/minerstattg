package templates

import (
	"bytes"
	"log"
	"text/template"

	"github.com/GeorgijGrigoriev/minerstattg/internal/api"
)

type MarketsTemplate struct {
	Exchange string
	Response api.Market
}

func MakeTemplate(exch string, resp api.Market) string {
	td := MarketsTemplate{Exchange: exch, Response: resp}

	t, err := template.New("MarketMsg").Parse("Exchange: {{ .Exchange}}\nPrice: {{ .Response.Price}}\nHigh: {{ .Response.High}}\nLow: {{ .Response.Low}}\nVolume: {{ .Response.Volume}}")
	if err != nil {
		log.Printf("templater error: %s", err.Error())

		return ""
	}

	tmpl := new(bytes.Buffer)

	t.Execute(tmpl, td)

	return tmpl.String()
}
