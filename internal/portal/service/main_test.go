package service

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"v2ray-saas-gemini/internal/config"
	"v2ray-saas-gemini/internal/database"
	"v2ray-saas-gemini/internal/models"
)

// TestMain runs before all tests in the package.
func TestMain(m *testing.M) {
	// Setup
	setup()
	// Run tests
	code := m.Run()
	// Teardown
	teardown()
	// Exit
	os.Exit(code)
}

func setup() {
	// Find the project root by looking for go.mod
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			break // Found the root
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			log.Fatalf("go.mod not found in any parent directory")
		}
		dir = parent
	}

	// Change to the project root directory
	if err := os.Chdir(dir); err != nil {
		log.Fatalf("Failed to change to root directory: %v", err)
	}

	// Load config for database connection
	config.Init()
	os.Setenv("DB_DSN", config.Cfg.Database.DSN)
	database.Init()
	log.Println("Test database setup complete.")
}

func teardown() {
	// Clean up the database after tests
	err := database.DB.Migrator().DropTable(
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
		log.Fatalf("Failed to drop tables: %v", err)
	}
	log.Println("Test database teardown complete.")
}