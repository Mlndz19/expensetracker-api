package dtos

import (
	"time"
)

type UpdateTransaction struct {
	Date             *time.Time `json:"date" binding:"omitempty"`
	Description      string     `json:"description" binding:"omitempty"`
	Notes            string     `json:"notes" binding:"omitempty"`
	Amount           *float64   `json:"amount" binding:"omitempty,gt=0"`
	PaymentMethodID  *uint      `json:"payment_method" binding:"omitempty"`
	CardID           *uint      `json:"card_id" binding:"omitempty"`
}