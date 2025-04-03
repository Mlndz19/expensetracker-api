package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `json:"email" binding:"required,email" gorm:"unique;not null"`
	Username string    `json:"username" binding:"required,min=5,max=25" gorm:"unique;not null"`
	Password string    `json:"password" binding:"required" gorm:"not null"`
	Active   bool      `json:"active" gorm:"default:true"`
}