package routes

import (
	"expensetrack/main.go/controllers"

	"github.com/gin-gonic/gin"
)

func BankRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/bank")
	{
		userGroup.GET("/", controllers.GetAllBanks)
		userGroup.GET("/:id", controllers.GetBankById)
		userGroup.POST("/", controllers.CreateBank)
		userGroup.PUT("/:id", controllers.UpdateBankById)
		userGroup.DELETE("/:id", controllers.DeleteBankById)
	}
}