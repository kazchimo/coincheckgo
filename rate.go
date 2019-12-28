package coincheck

import `encoding/json`

type Rate struct {
	client *CoinCheck
}

type ShopRateResponse struct {
	Rate float64 `json:",string"`
}

func (r Rate) Shop(pair CurrencyPair) (ShopRateResponse, error) {
	response := r.client.Request("GET", "api/rate/" + pair.pairString, "")
	var rateResponse ShopRateResponse
	err := json.Unmarshal([]byte(response), &rateResponse)

	return rateResponse, err
}




