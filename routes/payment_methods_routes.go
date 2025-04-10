package routes

import (
	"expensetrack/main.go/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentMethodsRoutes(r *gin.RouterGroup){
	routerGroup := r.Group("/payment-methods")
	{
		routerGroup.GET("/", controllers.GetAllPaymentMethods)
		routerGroup.GET("/:id", controllers.GetPaymentMethodById)
	}
}