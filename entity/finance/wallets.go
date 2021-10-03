package entity

import "github.com/google/uuid"

// Wallet contains the sum of all assets
type Wallet struct {
	UserID  uuid.UUID `json:"user_id"`
	Balance float64   `json:"balance"`
}