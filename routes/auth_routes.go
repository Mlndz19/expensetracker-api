package routes

import (
	"expensetrack/main.go/auth"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup){
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", auth.RegisterUser)
		authRoute.POST("/login", auth.LoginUser)
	}
}