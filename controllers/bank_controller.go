package controllers

import (
	"expensetrack/main.go/config"
	"expensetrack/main.go/models"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetAllBanks(c *gin.Context){
	var banks []models.Bank

	if err := config.DB.Find(&banks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"banks": banks})
}

func GetBankById(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}


	var bank models.Bank
	if err := config.DB.First(&bank, id).Error; err != nil{
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"bank": bank})
}

func CreateBank(c *gin.Context){
	var bank models.Bank

	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&bank).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Bank created correctly",
		"bank": bank,
	})
}

func UpdateBankById(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var bank models.Bank
	if err := config.DB.First(&bank, id).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input struct{
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Model(&bank).Updates(models.Bank{Name: input.Name}).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	bank.Name = input.Name

	c.JSON(http.StatusOK, gin.H{
		"message": "Bank updated correctly",
		"bank": bank,
	})
}

func DeleteBankById(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var bank models.Bank
	if err := config.DB.First(&bank, id).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound, gin.H{"error": "Bank not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&bank).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank deleted correctly"})
}