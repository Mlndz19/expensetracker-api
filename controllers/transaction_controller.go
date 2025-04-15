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

func UpdateTransactionByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var transaction models.Transaction
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transacción no encontrada para este usuario"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var dto dtos.UpdateTransaction
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const PaymentMethodCash = 1

	if dto.PaymentMethodID != nil {
		if *dto.PaymentMethodID == PaymentMethodCash && dto.CardID != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Transacciones en efectivo no deben tener tarjeta"})
			return
		}
		if *dto.PaymentMethodID != PaymentMethodCash && dto.CardID == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Este método de pago requiere una tarjeta"})
			return
		}
	}

	if dto.CardID != nil {
		var card models.Card
		if err := config.DB.Where("id = ? AND user_id = ?", *dto.CardID, userID).First(&card).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Tarjeta inválida o no pertenece al usuario"})
			return
		}
	}

	updateData := map[string]interface{}{}

	if dto.Description != "" {
		updateData["description"] = dto.Description
	}
	if dto.Notes != "" {
		updateData["notes"] = dto.Notes
	}
	if dto.Date != nil {
		updateData["date"] = *dto.Date
	}
	if dto.Amount != nil {
		updateData["amount"] = *dto.Amount
	}
	if dto.PaymentMethodID != nil {
		updateData["payment_method_id"] = *dto.PaymentMethodID
	}
	if dto.CardID != nil {
		updateData["card_id"] = *dto.CardID
	} else if dto.PaymentMethodID != nil && *dto.PaymentMethodID == PaymentMethodCash {
		updateData["card_id"] = nil
	}

	if err := config.DB.Model(&transaction).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la transacción"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transacción actualizada correctamente",
		"transaction": transaction,
	})
}

func DeleteTransactionByID(c *gin.Context){
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var transaction models.Transaction
	if err := config.DB.Where("id = ? AND user_id = ?", id, userID).First(&transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transacción no encontrada para este usuario"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Delete(&transaction).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted correctly"})
}