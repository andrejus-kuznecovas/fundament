package database

import (
	"log"
	"os"
	"time"

	"github.com/yourusername/fundament/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// Default connection string for development
		dsn = "postgresql://fundament:fundament123@postgres:5432/fundament"
	}

	log.Printf("Attempting to connect to database: %s", dsn)

	// Retry connection logic
	maxRetries := 10
	var db *gorm.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err == nil {
			break
		}

		log.Printf("Database connection attempt %d/%d failed: %v", i+1, maxRetries, err)

		if i < maxRetries-1 {
			// Wait before retrying (exponential backoff)
			waitTime := 1 << i // 1, 2, 4, 8, 16 seconds
			if waitTime > 30 {
				waitTime = 30
			}
			log.Printf("Retrying in %d seconds...", waitTime)
			time.Sleep(time.Duration(waitTime) * time.Second)
		}
	}

	if err != nil {
		log.Printf("❌ All database connection attempts failed: %v", err)
		log.Println("⚠️  Starting server without database connection")
		log.Println("💡 The application will attempt to reconnect on each request")
		return nil, nil // Allow app to start without database
	}

	// Get underlying sql.DB for connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Warning: Failed to get underlying sql.DB: %v", err)
		return db, nil
	}

	// Configure connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Test the connection
	if err := sqlDB.Ping(); err != nil {
		log.Printf("Warning: Database ping failed: %v", err)
		return db, nil
	}

	// Auto-migrate the schema
	if err := db.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
		log.Printf("Warning: Failed to migrate database: %v", err)
		log.Println("Continuing without database migration...")
		return db, nil
	}

	log.Println("✅ Database connected and migrated successfully")
	return db, nil
}
