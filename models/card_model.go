package models

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"not null"`
	Number string `json:"number" binding:"required,min=16,max=16" gorm:"unique;not null"`
	BankID uint `json:"bank_id" binding:"required"`
	Bank Bank `json:"bank" gorm:"foreignKey:BankID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	UserID uint `json:"user_id" binding:"required"`
	User User `json:"user" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
}