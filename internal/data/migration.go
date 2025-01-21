package data

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dsn string) *gorm.DB {
	// Initialize database
	// dsn := "host=" + cfg.DatabaseConfig.Host + " user=" + cfg.DatabaseConfig.User + " password=" + cfg.DatabaseConfig.Password + " dbname=" + cfg.DatabaseConfig.Name + " port=" + cfg.DatabaseConfig.Port + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s\n", err.Error())
	}

	return db
}
