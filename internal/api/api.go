package api

type API interface {
	GetMarket() error
}

func GetMarkets(a API) error {
	if err := a.GetMarket(); err != nil {
		return err
	}

	return nil
}
