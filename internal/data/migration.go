package data

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s\n", err.Error())
	}

	// err = RunMigrations(db)
	if err != nil {
		log.Fatalf("Failed to run migrations: %s\n", err.Error())
	}

	return db
}

// func RunMigrations(db *gorm.DB) error {
// 	return db.AutoMigrate(
// 		&schema.User{},
// 	)
// }
