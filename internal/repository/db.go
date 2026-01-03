package repository

import (
	"excel-processor/internal/entity"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSsl := os.Getenv("DB_SSLMODE")
	dbTime := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbPort, dbSsl, dbTime)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal connect database: ", err)
	}

	err = DB.AutoMigrate(&entity.Student{})
	if err != nil {
		log.Fatal("❌ Gagal migrasi database: ", err)
	}

	fmt.Println("🚀 Database Connected & Migrated!")
}
