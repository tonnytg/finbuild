package database

import (
	"database/sql"
	"errors"
	entity "finbuild/entity/finance"

	log "finbuild/pkg/logging"
	"fmt"
	"github.com/google/uuid"
)

type FinanceDb struct {
	db *sql.DB
}

func NewFinanceDb(db *sql.DB) *FinanceDb {
	return &FinanceDb{db: db}
}

func UpdateBalance(userid uuid.UUID, action string, quantity float64, price float64) (float64, error) {

	// get balance
	var account entity.Wallet
	DB.Table("accounts").First(&account).Scan(&account)

	if action == "BUY" {
		// update balance
		v := account.Balance + (price * quantity)
		DB.Table("accounts").Model(&entity.Wallet{}).Where("user_id = ?", userid).Update("balance", v)
		return v, nil
	}
	if action == "SELL" {
		// update balance
		v := account.Balance - (price * quantity)
		DB.Model(&entity.Wallet{}).Where("user_id = ?", userid).Update("balance", v)
		return v, nil
	}

	return 0, errors.New("wrong action")
}

func RegisterAsset(userid uuid.UUID, action string, code string, quantity float64, price float64, date string) error {
	msg := fmt.Sprintf("save in database - user: %s action: %s quantity: %f price: %f", userid, action, quantity, price)
	log.Msg("INFO", msg)
	as := &entity.Asset{
		userid,
		code,
		price,
		quantity,
		action,
		date,
	}
	DB.Model(&entity.Asset{}).Create(&as)
	return nil
}
