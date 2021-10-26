package exchange

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

type Interface interface {
	IsValid() bool
	GetID() string
}

func NewExchange() Exchanges {
	return Exchanges{}
}

func (e *Exchanges) IsValid() bool {

	if e.Quantity > 0 {
		return true
	}
	return false
}

func (e *Exchanges) GetID() string {
	return e.WalletID
}
