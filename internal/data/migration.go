package data

import (
	"finance-tracker/internal/auth/schema"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=finance-tracker port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s\n", err.Error())
	}

	return db
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&schema.User{},
	)
}
