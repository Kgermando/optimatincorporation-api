package database

import (
	"log"

	"github.com/kgermando/optimatincorporation-api/config"
	"github.com/kgermando/optimatincorporation-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *config.Config) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Portfolio{},
		&models.Service{},
		&models.Blog{},
		&models.TeamMember{},
		&models.Contact{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	DB = db
	log.Println("Database connected and migrated")
}
