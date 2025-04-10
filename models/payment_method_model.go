package models

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Description string `json:"description" binding:"required" gorm:"unique;not null"`
}

const (
	PaymentMethodCash = "Efectivo"
	PaymentMethodDebit = "Tarjeta de débito"
	PaymentMethodCredit = "Tarjeta de crédito"
)