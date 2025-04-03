package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Name string `json:"name" binding:"required" gorm:"unique;not null"`
}