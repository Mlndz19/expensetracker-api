package config

import (
	"fmt"
	"log"
	"os"

	"expensetrack/main.go/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar variables de entorno: ", err.Error())
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar con la base de datos: ", err.Error())
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("❌ Error al migrar la tabla User: ", err.Error())
	} else {
		fmt.Println("✔ Migración de la tabla User completada")
	}

	fmt.Println("✔ Conexión exitosa a postgres")
}