package controllers

import (
	"expensetrack/main.go/config"
	"expensetrack/main.go/models"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetAllPaymentMethods(c *gin.Context){
	var payment_methods []models.PaymentMethod
	if err := config.DB.Find(&payment_methods).Error; err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment_methods": payment_methods})
}

func GetPaymentMethodById(c *gin.Context){
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var payment_method models.PaymentMethod
	if err := config.DB.First(&payment_method, id).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment method not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payment_method": payment_method})
}