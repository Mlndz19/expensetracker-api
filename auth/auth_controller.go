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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
		return
	}

	if !CheckPassword(user.Password, loginData.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciales inválidas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login correcto"})

}

func RegisterUser(c *gin.Context){
	var newUSer models.User

	if err := c.ShouldBindJSON(&newUSer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := HashPassword(newUSer.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newUSer.Password = hashedPassword

	if err := config.DB.Create(&newUSer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created correctly",
		"user": newUSer,
	})
}