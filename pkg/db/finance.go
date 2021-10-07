package db

import (
	"errors"
	"finbuild/entity/finance"
	"github.com/google/uuid"

	log "finbuild/pkg/logging"
	"fmt"
)

func UpdateBalance(f *finance.Exchanges) (float64, error) {

	// get balance
	var account finance.Wallet
	DB.Table("wallets").First(&account).Scan(&account)

	if f.Action == "BUY" {
		// update balance
		v := account.Balance + (f.Price * f.Quantity)
		DB.Table("wallets").Model(&finance.Wallet{}).Where("wallet_id = ?", f.WalletID).Update("balance", v)
		return v, nil
	}
	if f.Action == "SELL" {
		// update balance
		v := account.Balance - (f.Price * f.Quantity)
		DB.Table("wallets").Model(&finance.Wallet{}).Where("wallet_id = ?", f.WalletID).Update("balance", v)
		return v, nil
	}

	return 0, errors.New("wrong action")
}

func RegisterExchange(f *finance.Exchanges) error {

	if f.Quantity > 0 {
		msg := fmt.Sprintf("save in db - wallet_id: %s asset_id: %s action: %s quantity: %f price: %f",
			f.WalletID, f.AssetID, f.Action, f.Quantity, f.Price)

		log.Msg("INFO", msg)

		DB.Table("exchanges").Model(&finance.Exchanges{}).Create(&f)
	}
	return nil
}

func GetWalletsByUser(s *finance.ShortWallet, u uuid.UUID) *finance.ShortWallet{
	DB.Table("wallets").First(&s, "user_id = ?", u).Scan(&s).Find(&finance.Wallet{})
	return s
}

func GetWalletsByID(u uuid.UUID) *finance.ShortWallet{
	var s *finance.ShortWallet
	DB.Table("wallets").First(&s, "wallet_id = ?", u).Scan(&s).Find(&finance.Wallet{})
	return s
}
