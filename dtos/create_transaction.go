package dtos

type CreateTransaction struct {
	Description string `json:"description" binding:"required,min=1,max=50"`
	Notes string `json:"notes" binding:"max=150"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
	PaymentMethodID uint `json:"payment_method" binding:"required"`
	CardID *uint `json:"card_id"`
}