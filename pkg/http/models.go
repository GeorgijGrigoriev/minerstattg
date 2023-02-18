package http

type TradeOgreAPIMarkets struct {
	Success      string `json:"success"`
	InitialPrice string `json:"initialprice"`
	Price        string `json:"price"`
	High         string `json:"high"`
	Low          string `json:"low"`
	Volume       string `json:"volume"`
	Bid          string `json:"bid"`
	Ask          string `json:"ask"`
}

type BybitAPIMarkets struct {
	RetCode int                   `json:"retCode"`
	RetMsg  string                `json:"retMsg"`
	Result  BybitAPIMarketsResult `json:"result"`
}

type BybitAPIMarketsResult struct {
	Category string                      `json:"category"`
	List     []BybitAPIMarketsResultList `json:"list"`
}

type BybitAPIMarketsResultList struct {
	LastPrice string `json:"lastPrice"`
	High      string `json:"highPrice24h"`
	Low       string `json:"lowPrice24h"`
	Volume    string `json:"volume24h"`
}
