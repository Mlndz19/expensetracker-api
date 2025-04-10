package main

import (
	"fmt"
	"log"
	"os"

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
	{
		routes.UsersRoutes(api)
		routes.AuthRoutes(api)
		routes.BankRoutes(api)
		routes.PaymentMethodsRoutes(api)
	}

	SERVER_PORT := os.Getenv("SERVER_PORT")
	port := fmt.Sprintf(":%s", SERVER_PORT)
	r.Run(port)
}