package main

import (
	"fmt"
	"log"
	"net/http"
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
	routes.RegisterUsersRoutes(r);


	r.GET("api/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	SERVER_PORT := os.Getenv("SERVER_PORT")
	port := fmt.Sprintf(":%s", SERVER_PORT)
	r.Run(port)
}