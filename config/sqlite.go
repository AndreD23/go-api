package config

import (
	"github.com/AndreD23/go-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	// Check if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found. Creating...")

		// Create database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Errorf("cannot create directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("cannot create database file: %v", err)
			return nil, err
		}

		err = file.Close()
		if err != nil {
			logger.Errorf("cannot close database file: %v", err)
			return nil, err
		}
	}

	// Create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("error Initialize SQLite: %v", err)
		return nil, err
	}

	// Migrate the schemas
	err = db.AutoMigrate(&schemas.Opportunity{})
	if err != nil {
		logger.Errorf("error auto migrate database: %v", err)
		return nil, err
	}

	return db, nil
}
