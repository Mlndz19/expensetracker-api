package dtos

type UpdateCard struct {
	Name string `json:"name" binding:"omitempty"`
	Number string `json:"number" binding:"omitempty,min=16,max=16"`
	BankID uint `json:"bank_id" binding:"omitempty"`
}