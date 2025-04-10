package config

import (
	"expensetrack/main.go/models"
	"fmt"
)

func SeedPaymentMethods() {
	methods := []string{
		models.PaymentMethodCash,
		models.PaymentMethodDebit,
		models.PaymentMethodCredit,
	}

	for _, description := range methods{
		var existing models.PaymentMethod
		err := DB.Where("description = ?", description).First(&existing).Error
		if err != nil{
			if err := DB.Create(&models.PaymentMethod{Description: description}).Error; err != nil{
				fmt.Println("❌ Error al insertar método de pago:", description, "-", err.Error())
			} else {
				fmt.Println("✅ Método de pago insertado:", description)
			}
		}
	}
}