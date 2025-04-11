package routes

import (
	"expensetrack/main.go/controllers"

	"github.com/gin-gonic/gin"
)

func CardRoutes(r *gin.RouterGroup){
	cardRoutes := r.Group("/cards")
	{
		cardRoutes.GET("/", controllers.GetAllCardsByUserID)
		cardRoutes.GET("/:id", controllers.GetCardByID)
		cardRoutes.POST("/", controllers.CreateCard)
	}
}