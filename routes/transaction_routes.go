package routes

import (
	"expensetrack/main.go/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(r *gin.RouterGroup) {
	transactionRoutes := r.Group("/transactions")
	{
		transactionRoutes.GET("/", controllers.GetTransacionsByUserID)
		transactionRoutes.GET("/:id", controllers.GetTransacionByID)
	}
}