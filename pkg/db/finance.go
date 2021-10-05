package db

import (
	"errors"
	"finbuild/entity/finance"
	"github.com/google/uuid"

	log "finbuild/pkg/logging"
	"fmt"
)

func UpdateBalance(WalletID string, Action string, Quantity float64, Price float64, Tax float64, Date string) (float64, error) {

	// get balance
	var account finance.Wallet
	DB.Table("wallets").First(&account).Scan(&account)

	if Action == "BUY" {
		// update balance
		v := account.Balance + (Price * Quantity)
		DB.Table("wallets").Model(&finance.Wallet{}).Where("wallet_id = ?", WalletID).Update("balance", v)
		return v, nil
	}
	if Action == "SELL" {
		// update balance
		v := account.Balance - (Price * Quantity)
		DB.Table("wallets").Model(&finance.Wallet{}).Where("wallet_id = ?", WalletID).Update("balance", v)
		return v, nil
	}

	return 0, errors.New("wrong action")
}

func RegisterExchange(WalletID string, AssetID string, Action string, Tax float64, Quantity float64, Price float64, Date string) error {
	msg := fmt.Sprintf("save in db - wallet_id: %s asset_id: %s action: %s quantity: %f price: %f", WalletID, AssetID, Action, Quantity, Price)
	log.Msg("INFO", msg)
	as := &finance.Exchanges{
		WalletID: WalletID,
		AssetID:  AssetID,
		Price:    Price,
		Tax:      Tax,
		Quantity: Quantity,
		Action:   Action,
		Date:     Date,
	}
	DB.Table("exchanges").Model(&finance.Exchanges{}).Create(&as)
	return nil
}

func GetWallets(shortWallet *finance.ShortWallet, u uuid.UUID) *finance.ShortWallet{
	DB.Table("wallets").First(&shortWallet, "user_id = ?", u).Scan(&shortWallet).Find(&finance.Wallet{})
	return shortWallet
}