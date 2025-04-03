package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Guid     string    `json:"guid" gorm:"unique;not null"`
	Email    string    `json:"email" binding:"required,email" gorm:"unique;not null"`
	Username string    `json:"username" binding:"required,min=5,max=25" gorm:"unique;not null"`
	Password string    `json:"-" binding:"required" gorm:"not null"`
	Active   bool      `json:"active" gorm:"default:true"`
}