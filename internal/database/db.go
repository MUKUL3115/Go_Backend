package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-backend/internal/models"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	// Load .env if present (optional)
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Connect using GORM with the postgres driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Supabase DB: %v", err)
	}

	// Example: Auto-migrate your models
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	DB = db
	return DB
}
