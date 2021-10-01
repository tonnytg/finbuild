package database

import (
	"database/sql"
)

type FinanceDb struct {
	db *sql.DB
}

func NewFinanceDb(db *sql.DB) *FinanceDb {
	return &FinanceDb{db: db}
}

