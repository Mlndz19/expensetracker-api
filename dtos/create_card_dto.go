package dtos

type CreateCard struct {
	Name string `json:"name" binding:"required"`
	Number string `json:"number" binding:"required,min=16,max=16"`
	BankID uint `json:"bank_id" binding:"required"`
}