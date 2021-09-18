package finance

// Asset it's a term for all kind of item in Finance Market
type Asset struct {
	ID       string
	Price    float64
	Quantity float64
	Date     string
	Action   string // BUY or SELL
}

// AssetRent it's all transaction/action gain or lost money
type AssetRent struct {
	Date    string
	Price   float64
	AssetID string
	Type    string
	Percent float64
}

// Account contains the sum of all assets
type Account struct {
	Balance float64
}
