package data

import (
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
	// &schema.User{},
	// &models.Order{},
	)
}
