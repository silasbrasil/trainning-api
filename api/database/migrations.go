package database

import (
	"trainnig-api-poc/api/users/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
