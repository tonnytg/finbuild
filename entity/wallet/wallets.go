package wallet

import "github.com/google/uuid"

// Wallet contains the sum of all assets
type Wallet struct {
	WalletID uuid.UUID `json:"wallet_id"`
	UserID   uuid.UUID `json:"user_id"`
	Balance  float64   `json:"balance"`
}

type ShortWallet struct {
	WalletID string
	Balance  float64
}

// NewWallet create a new Wallet with 0 balance
func NewWallet(u uuid.UUID) *Wallet {
	wallet := Wallet{
		WalletID: uuid.New(),
		UserID:   u,
		Balance:  0,
	}
	return &wallet
}

// Rent it's all transaction/action gain or lost money
// contains information about history of wallet
type Rent struct {
	WalletID uuid.UUID `json:"wallet_id"`
	Price    float64   `json:"price"`
	AssetID  string    `json:"assetId"`
	Type     string    `json:"type"`
	Percent  float64   `json:"percent"`
	Date     string    `json:"date"`
}

type Interface interface {
	IsValid(w *Wallet) bool
}

func (w *Wallet) IsValid() bool {
	return true
}
