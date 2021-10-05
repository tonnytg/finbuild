package entity

// Exchanges it's a term for all kind of item in Finance Market
// contains registry about transaction purchase or sellers orders
type Exchanges struct {
	WalletID string  `json:"wallet_id"`
	AssetID  string  `json:"asset_id"`
	Price    float64 `json:"price"`
	Tax      float64 `json:"tax"`
	Quantity float64 `json:"quantity"`
	Action   string  `json:"action"` // BUY or SELL
	Date     string  `json:"date"`
}
