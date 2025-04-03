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

	MigrateDB()

	fmt.Println("✔ Conexión exitosa a postgres")
}

func MigrateDB(){
	err := DB.AutoMigrate(&models.User{}, &models.Bank{}, models.PaymentMethod{}, models.Card{}, models.Transaction{})
	if err != nil {
		log.Fatal("❌ Error al migrar la tabla User: ", err.Error())
	} else {
		fmt.Println("✔ Migración de la tabla User completada")
	}

	DB.Exec(`
    DO $$
    BEGIN
        IF NOT EXISTS (
            SELECT 1 FROM pg_constraint WHERE conname = 'check_card_based_on_payment_method'
        ) THEN
            ALTER TABLE transactions
            ADD CONSTRAINT check_card_based_on_payment_method
            CHECK (
                (payment_method_id = 1 AND card_id IS NULL) OR
                (payment_method_id != 1 AND card_id IS NOT NULL)
            );
        END IF;
    END$$;
	`)
}