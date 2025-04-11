package routes

import (
	"expensetrack/main.go/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(r *gin.RouterGroup) {
	userGroup := r.Group("/users")
	{
		userGroup.GET("/", controllers.GetAllUsers)
		userGroup.GET("/:id", controllers.GetUserById)
		userGroup.PUT("/:id", controllers.UpdateUserById)
		userGroup.DELETE("/:id", controllers.DeleteUserById)
	}
	
}