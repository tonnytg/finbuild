package entity

import (
	"github.com/google/uuid"
)

// Asset it's a term for all kind of item in Finance Market
// contains registry about transaction purchase or sellers orders
type Exchanges struct {
	UserID   uuid.UUID `json:"user_id"`
	ID       string    `json:"id"`
	Price    float64   `json:"price"`
	Quantity float64   `json:"quantity"`
	Action   string    `json:"action"` // BUY or SELL
	Date     string    `json:"date"`
}
