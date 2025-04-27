package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Description string `json:"description" binding:"required" gorm:"not null"`
}