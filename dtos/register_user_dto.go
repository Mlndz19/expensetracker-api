package dtos

type RegisterUser struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=5,max=25"`
	Password string `json:"password" binding:"required"`
}