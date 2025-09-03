package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	connErr error
)

func Init() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	url := os.Getenv("DATABASE_URL")
	db, connErr = gorm.Open(postgres.Open(url), &gorm.Config{})
	if connErr != nil {
		return nil, connErr
	}

	return db, nil
}

func GetDb() *gorm.DB {
	if db != nil {
		return db
	}

	log.Printf("Database not initialized. Call Init() first.")

	return nil
}
