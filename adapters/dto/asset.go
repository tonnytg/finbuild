package dto

import (
	"github.com/google/uuid"
)

type Asset struct {
	UserID   uuid.UUID `json:"user_id"`
	ID       string    `json:"id"`
	Price    float64   `json:"price"`
	Quantity float64   `json:"quantity"`
	Action   string    `json:"action"`
	Date     string    `json:"date"`
}

func NewAsset() *Asset {
	return &Asset{}
}
