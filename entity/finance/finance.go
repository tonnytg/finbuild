package finance

import (
	"errors"
	"finbuild/database"
	"github.com/google/uuid"
	"log"
)

type AssetInterface interface {
	UpdateBalance(userid uuid.UUID, action string, quantity float64,price float64) (float64, error)
	RegisterAsset(userid uuid.UUID, action string, quantity float64,price float64) error
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

func (a *Account) UpdateBalance(userid uuid.UUID, action string, quantity float64,price float64) (float64, error){

	// get balance
	var account Account
	database.DB.Table("accounts").First(&account).Scan(&account)

	if action == "BUY" {
		// update balance
		v := account.Balance + ( price * quantity )
		database.DB.Model(&Account{}).Where("user_id = ?", userid).Update("balance", v)
		return v, nil
	}
	if action == "SELL" {
		// update balance
		v := account.Balance - ( price * quantity )
		database.DB.Model(&Account{}).Where("user_id = ?", userid).Update("balance", v)
		return v, nil
	}

	return 0, errors.New("wrong action")
}

func (a *Account) RegisterAsset(userid uuid.UUID, action string, quantity float64,price float64) error {
	log.Printf("INFO: save in database - user: %s action: %s quantity: %f price: %f", userid, action, quantity, price)
	return nil
}
