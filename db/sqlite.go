package db

import (
	"bitbucket.org/guardrails-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("repos_local.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.Repository{})
	db.AutoMigrate(&models.ScanResults{})
	return db, nil
}
