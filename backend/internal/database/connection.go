package database

import (
	"log"
	"os"

	"github.com/yourusername/fundament/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Default connection string for development
		dsn = "postgresql://fundament:fundament123@localhost:5432/fundament"
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("Warning: Database connection failed: %v", err)
		log.Println("Continuing without database connection for development...")
		return nil, nil // Return nil instead of error to allow app to start
	}

	// Get underlying sql.DB for connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Warning: Failed to get underlying sql.DB: %v", err)
		return nil, nil
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Auto-migrate the schema
	if err := db.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
		log.Printf("Warning: Failed to migrate database: %v", err)
		log.Println("Continuing without database migration...")
		return db, nil
	}

	log.Println("Database connected and migrated successfully")
	return db, nil
}
