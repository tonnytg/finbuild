package database

import (
	"database/sql"
	"finbuild/entity/finance"
)

type FinanceDb struct {
	db *sql.DB
}

func NewFinanceDb(db *sql.DB) *FinanceDb {
	return &FinanceDb{db: db}
}

func (f *FinanceDb) Get(id string) (finance.AssetInterface, error) {
	var asset finance.Asset
	stmt, err := f.db.Prepare("select user_id, id, price, quantity, action, date from assets where user_id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&asset.UserID, &asset.ID, &asset.Price, &asset.Quantity, &asset.Action, &asset.Date)
	if err != nil {
		return nil, err
	}
	return &asset, nil
}
