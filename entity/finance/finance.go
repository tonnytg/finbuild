package finance

import (
	"finbuild/database"
	"github.com/google/uuid"
)

type AssetInterface interface {
	UpdateBalance() (float64, error)
}

// Asset it's a term for all kind of item in Finance Market
type Asset struct {
	UserID   uuid.UUID `json:"user_id"`
	ID       string    `json:"id"`
	Price    float64   `json:"price"`
	Quantity float64   `json:"quantity"`
	Action   string    `json:"action"` // BUY or SELL
	Date     string    `json:"date"`
}

// AssetRent it's all transaction/action gain or lost money
type AssetRent struct {
	UserID  uuid.UUID `json:"user_id"`
	Date    string    `json:"date"`
	Price   float64   `json:"price"`
	AssetID string    `json:"assetId"`
	Type    string    `json:"type"`
	Percent float64   `json:"percent"`
}

// Account contains the sum of all assets
type Account struct {
	UserID  uuid.UUID `json:"user_id"`
	Balance float64   `json:"balance"`
}

func (a *Account) UpdateBalance(userid uuid.UUID, price float64) (float64, error){

	// get balance
	var account Account
	database.DB.Table("accounts").First(&account).Scan(&account)

	// update balance
	v := account.Balance + price
	database.DB.Model(&Account{}).Where("user_id = ?", userid).Update("balance", v)

	return v, nil
}