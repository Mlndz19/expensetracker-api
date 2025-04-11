package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Date time.Time `json:"date" gorm:"type:date;not null"`
	Description string `json:"description" binding:"required" gorm:"not null"`
	Notes string `json:"notes"`
	Amount float64 `json:"amount" binding:"required,gt=0" gorm:"not null"`
	PaymentMethodID uint `json:"payment_method_id" binding:"required" gorm:"not null"`
	PaymentMethod PaymentMethod `json:"payment_method" gorm:"foreignKey:PaymentMethodID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	CardID *uint `json:"card_id" gorm:"default:null"`
	Card Card `json:"card" gorm:"foreignKey:CardID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserID uint `json:"user_id" binding:"required" gorm:"not null"`
	User User `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}