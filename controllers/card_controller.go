package controllers

import (
	"net/http"
	"strconv"

	"expensetrack/main.go/config"
	"expensetrack/main.go/dtos"
	"expensetrack/main.go/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetCardByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var card models.Card
	if err := config.DB.Where("id = ? and user_id = ?", id, userID).First(&card).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"card": card})
}

func GetAllCardsByUserID(c *gin.Context){
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var cards []models.Card
	if err := config.DB.Preload("Bank").Where("user_id = ?", userID.(uint)).Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cards": cards})
}

func CreateCard(c *gin.Context){
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var newCardInput dtos.CreateCard
	if err := c.ShouldBindJSON(&newCardInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCard := models.Card{
		Name: newCardInput.Name,
		Number: newCardInput.Number,
		BankID: newCardInput.BankID,
		UserID: userID.(uint),
	}

	if err := config.DB.Create(&newCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Card created correctly",
		"card": newCard,
	})

}

func UpdateCardByID(c *gin.Context){
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var card models.Card
	if err := config.DB.Where("id = ? and user_id = ?", id, userID).First(&card).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found for this user"})
		return
	}

	var updatedCard dtos.UpdateCard
	if err := c.ShouldBindJSON(&updatedCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cardData := map[string]interface{}{}
	if updatedCard.Name != "" {
		cardData["name"] = updatedCard.Name
	}

	if updatedCard.Number != "" {
		cardData["number"] = updatedCard.Number
	}

	if err := config.DB.Model(&card).Updates(cardData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Card updated correctly",
		"card": card,
	})
}

func DeleteCardByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var card models.Card
	if err := config.DB.Where("id = ? and user_id = ?", id, userID).First(&card).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found for this user"})
		return
	}

	if err := config.DB.Delete(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Card deleted correctly"})
}