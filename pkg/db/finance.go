package db

import (
	"errors"
	"finbuild/entity/exchange"
	"finbuild/entity/wallet"

	"github.com/google/uuid"

	log "finbuild/pkg/logging"
	"fmt"
)

func UpdateBalance(f *exchange.Exchanges) (float64, error) {

	// get balance
	var w wallet.Wallet
	DB.Table("wallets").First(&w, "wallet_id = ?", f.WalletID).Scan(&w)

	if f.Action == "BUY" {
		// update balance
		v := w.Balance + (f.Price * f.Quantity)
		DB.Table("wallets").Model(&wallet.Wallet{}).Where("wallet_id = ?", f.WalletID).Update("balance", v)
		return v, nil
	}
	if f.Action == "SELL" {
		// update balance
		v := w.Balance - (f.Price * f.Quantity)
		DB.Table("wallets").Model(&wallet.Wallet{}).Where("wallet_id = ?", f.WalletID).Update("balance", v)
		return v, nil
	}

	return 0, errors.New("wrong action")
}

func RegisterExchange(e *exchange.Exchanges) error {

	if e.IsValid() {
		msg := fmt.Sprintf("save in db - wallet_id: %s asset_id: %s action: %s quantity: %e price: %e",
			e.WalletID, e.AssetID, e.Action, e.Quantity, e.Price)

		log.Msg("INFO", msg)

		DB.Table("exchanges").Model(&exchange.Exchanges{}).Create(&e)
	}
	return nil
}

func GetWalletsByUser(ws *wallet.ShortWallet, u uuid.UUID) *wallet.ShortWallet {
	DB.Table("wallets").First(&ws, "user_id = ?", u).Scan(&ws).Find(&wallet.Wallet{})
	return ws
}

func GetWalletsByID(u uuid.UUID) *wallet.ShortWallet {
	var ws *wallet.ShortWallet
	DB.Table("wallets").First(&ws, "wallet_id = ?", u).Scan(&ws).Find(&wallet.Wallet{})
	return ws
}
