package database

import (
	"fmt"
	"log"
	"os"
	"v2ray-saas-gemini/internal/config"
	"v2ray-saas-gemini/internal/models"
	"v2ray-saas-gemini/internal/utils"

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
		&models.RechargePreset{},
		&models.PasswordResetToken{},
		&models.Ticket{},
		&models.TicketReply{},
		&models.Announcement{},
		&models.HelpDocument{},
		&models.Admin{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}

	fmt.Println("Database schema migrated successfully.")

	// Create default admin if not exists
	createDefaultAdmin()
}

// createDefaultAdmin creates a default admin account if it doesn't exist
func createDefaultAdmin() {
	// Check if admin configuration is provided
	if config.Cfg.Admin.Username == "" || config.Cfg.Admin.Password == "" {
		log.Println("Admin configuration not provided, skipping default admin creation")
		return
	}

	// Check if admin already exists
	var existingAdmin models.Admin
	if err := DB.Where("username = ?", config.Cfg.Admin.Username).First(&existingAdmin).Error; err == nil {
		log.Printf("Admin user '%s' already exists, skipping creation", config.Cfg.Admin.Username)
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Printf("Error checking for existing admin: %v", err)
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(config.Cfg.Admin.Password)
	if err != nil {
		log.Printf("Failed to hash admin password: %v", err)
		return
	}

	// Set default role if not provided
	role := config.Cfg.Admin.Role
	if role == "" {
		role = "admin"
	}

	// Create admin account
	admin := models.Admin{
		Username:     config.Cfg.Admin.Username,
		PasswordHash: hashedPassword,
		Role:         role,
		IsActive:     true,
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Printf("Failed to create default admin: %v", err)
		return
	}

	log.Printf("Default admin user '%s' created successfully", config.Cfg.Admin.Username)
}
