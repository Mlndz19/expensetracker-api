package auth

import (
	"expensetrack/main.go/config"
	"expensetrack/main.go/dtos"
	"expensetrack/main.go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var loginData dtos.LoginModel

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciales inválidas"})
		return
	}

	if !CheckPassword(user.Password, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciales inválidas"})
		return
	}

	token, err := GenerateSignedToken(user.ID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login correcto",
		"token": token,
	})

}

func RegisterUser(c *gin.Context){
	var input dtos.RegisterUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var newUser models.User
	newUser.Email = input.Email
	newUser.Username = input.Username
	newUser.Password = hashedPassword

	if err := config.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created correctly",
		"user": newUser,
	})
}