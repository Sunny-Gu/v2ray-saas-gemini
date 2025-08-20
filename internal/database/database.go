package database

import (
	"fmt"
	"log"
	"os"
	"v2ray-saas-gemini/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init initializes the database connection and migrates the schema.
func Init() {
	var err error
	// DSN (Data Source Name) for the database connection.
	// It's recommended to use environment variables for sensitive data.
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable not set. Please configure it in config.yaml.")
	}

	// Open a connection to the database.
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established successfully.")

	// Auto-migrate the schema.
	// This will create the tables based on the GORM models.
	err = DB.AutoMigrate(
		&models.User{},
		&models.Balance{},
		&models.Subscription{},
		&models.RechargeOrder{},
		&models.SubscriptionOrder{},
		&models.Node{},
		&models.NodeTraffic{},
		&models.Plan{},
		&models.Coupon{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	fmt.Println("Database schema migrated successfully.")
}
