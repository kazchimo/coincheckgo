package coincheck

type Account struct {
	client *CoinCheck
}

type Balance struct {
	JPY string
	BTC string
}

type BalanceResponse struct {
	Success bool `json:"success"`
	Balance
	Error string
}

// Make sure a balance.
func (a Account) Balance() string {
	return a.client.Request("GET", "api/accounts/balance", "")
}

// Make sure a leverage balance.
func (a Account) LeverageBalance() string {
	return a.client.Request("GET", "api/accounts/leverage_balance", "")
}

// Get account information.
func (a Account) Info() string {
	return a.client.Request("GET", "api/accounts/balance", "")
}
