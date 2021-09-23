package finance

// Asset it's a term for all kind of item in Finance Market
type Asset struct {
	ID       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
	Action   string  `json:"action"` // BUY or SELL
	Date     string  `json:"date"`
}

// AssetRent it's all transaction/action gain or lost money
type AssetRent struct {
	Date    string  `json:"date"`
	Price   float64 `json:"price"`
	AssetID string  `json:"assetId"`
	Type    string  `json:"type"`
	Percent float64 `json:"percent"`
}

// Account contains the sum of all assets
type Account struct {
	Balance float64 `json:"balance"`
}