package controllers

import (
	"net/http"
	"strconv"
	"time"

	"expensetrack/main.go/config"
	"expensetrack/main.go/dtos"
	"expensetrack/main.go/models"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetTransacionByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var transaction models.Transaction
	if err := config.DB.Preload("Card").Preload("Payment_Method").Where("id = ? and user_id = ?", id, userID).First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found for this user"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}

func GetTransacionsByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var transactions []models.Transaction
	if err := config.DB.Find(&transactions, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

func CreateTransaction(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var transactionDto dtos.CreateTransaction
	if err := c.ShouldBindJSON(&transactionDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const PaymentMethodCash = 1

	if PaymentMethodCash == transactionDto.PaymentMethodID && transactionDto.CardID != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A transaction with cash payment method can not have a card"})
		return
	}

	if PaymentMethodCash != transactionDto.PaymentMethodID && transactionDto.CardID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This payment method must have a card"})
		return
	}

	if transactionDto.CardID != nil {
		var card models.Card
		if err := config.DB.Where("id = ? and user_id = ?", transactionDto.CardID, userID).First(&card).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Card not found for this user"})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	newTransaction := models.Transaction{
		Date: time.Now(),
		Description: transactionDto.Description,
		Notes: transactionDto.Notes,
		Amount: transactionDto.Amount,
		PaymentMethodID: transactionDto.PaymentMethodID,
		UserID: userID.(uint),
	}

	if transactionDto.CardID != nil {
		newTransaction.CardID = transactionDto.CardID
	}

	if err := config.DB.Create(&newTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transaction created correctly",
		"transaction": newTransaction,
	})

}