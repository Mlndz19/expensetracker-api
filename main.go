package main

import (
	"fmt"
	"log"
	"os"

	"expensetrack/main.go/auth"
	"expensetrack/main.go/config"
	"expensetrack/main.go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error al cargar el archivo .env: ", err.Error())
	}
	
	config.ConnectDB()

	r := gin.Default()

	api := r.Group("/api")
	routes.AuthRoutes(api)

	protected := api.Group("/")
	protected.Use(auth.JWTMiddleware())
	{
		routes.BankRoutes(protected)
		routes.PaymentMethodsRoutes(protected)
		routes.UsersRoutes(protected)
	}

	SERVER_PORT := os.Getenv("SERVER_PORT")
	port := fmt.Sprintf(":%s", SERVER_PORT)
	r.Run(port)
}